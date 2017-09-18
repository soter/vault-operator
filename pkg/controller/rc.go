package controller

import (
	"fmt"

	"github.com/appscode/go/log"
	stringz "github.com/appscode/go/strings"
	v1u "github.com/appscode/kutil/core/v1"
	"github.com/golang/glog"
	"github.com/hashicorp/vault/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/pkg/api/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

func (c *VaultController) runRCWatcher() {
	for c.processNextRC() {
	}
}

func (c *VaultController) processNextRC() bool {
	// Wait until there is a new item in the working queue
	key, quit := c.rcQueue.Get()
	if quit {
		return false
	}
	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
	// This allows safe parallel processing because two deployments with the same key are never processed in
	// parallel.
	defer c.rcQueue.Done(key)

	// Invoke the method containing the business logic
	err := c.runRCInitializer(key.(string))
	if err == nil {
		// Forget about the #AddRateLimited history of the key on every successful synchronization.
		// This ensures that future processing of updates for this key is not delayed because of
		// an outdated error history.
		c.rcQueue.Forget(key)
		return true
	}
	log.Errorln("Failed to process ReplicationController %v. Reason: %s", key, err)

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.rcQueue.NumRequeues(key) < c.options.MaxNumRequeues {
		glog.Infof("Error syncing deployment %v: %v", key, err)

		// Re-enqueue the key rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the key will be processed later again.
		c.rcQueue.AddRateLimited(key)
		return true
	}

	c.rcQueue.Forget(key)
	// Report to an external entity that, even after several retries, we could not successfully process this key
	runtime.HandleError(err)
	glog.Infof("Dropping deployment %q out of the queue: %v", key, err)
	return true
}

// syncToStdout is the business logic of the controller. In this controller it simply prints
// information about the deployment to stdout. In case an error happened, it has to simply return the error.
// The retry logic should not be part of the business logic.
func (c *VaultController) runRCInitializer(key string) error {
	obj, exists, err := c.rcIndexer.GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a ReplicationController, so that we will see a delete for one d
		fmt.Printf("ReplicationController %s does not exist anymore\n", key)
	} else {
		dp := obj.(*apiv1.ReplicationController)
		fmt.Printf("Sync/Add/Update for ReplicationController %s\n", dp.GetName())

		if dp.DeletionTimestamp != nil {
			if v1u.HasFinalizer(dp.ObjectMeta, "finalizer.kubernetes.io/vault") ||
				v1u.HasFinalizer(dp.ObjectMeta, "initializer.kubernetes.io/vault") {
				dp, err = v1u.PatchRC(c.k8sClient, dp, func(in *apiv1.ReplicationController) *apiv1.ReplicationController {
					in.ObjectMeta = v1u.RemoveFinalizer(in.ObjectMeta, "finalizer.kubernetes.io/vault")
					return in
				})
				return err
			}
		} else if dp.GetInitializers() != nil {
			pendingInitializers := dp.GetInitializers().Pending
			if pendingInitializers[0].Name == "vault.initializer.kubernetes.io" {
				serviceAccountName := stringz.Val(dp.Spec.Template.Spec.ServiceAccountName, "default")

				sa, err := c.k8sClient.CoreV1().ServiceAccounts(dp.Namespace).Get(serviceAccountName, metav1.GetOptions{})
				if err != nil {
					return err
				}

				var vaultSecret *v1.Secret
				if secretName, found := GetString(sa.Annotations, "vaultproject.io/secret.name"); !found {
					return fmt.Errorf("missing vault secret annotation for service account %s", serviceAccountName)
				} else {
					vaultSecret, err = c.k8sClient.CoreV1().Secrets(dp.Namespace).Get(secretName, metav1.GetOptions{})
					if err != nil {
						return err
					}
				}

				dp, err = v1u.PatchRC(c.k8sClient, dp, func(in *apiv1.ReplicationController) *apiv1.ReplicationController {
					in.ObjectMeta = v1u.RemoveNextInitializer(in.ObjectMeta)
					in.ObjectMeta = v1u.AddFinalizer(in.ObjectMeta, "finalizer.kubernetes.io/vault")

					volSrc := v1.SecretVolumeSource{
						SecretName: vaultSecret.Name,
						Items: []v1.KeyToPath{
							{
								Key:  api.EnvVaultAddress,
								Path: "vault-addr",
								// Mode:
							},
							{
								Key:  api.EnvVaultToken,
								Path: "token",
								// Mode:
							},
							{
								Key:  "VAULT_TOKEN_ACCESSOR",
								Path: "token-accessor",
								// Mode:
							},
							{
								Key:  "LEASE_DURATION",
								Path: "lease-duration",
								// Mode:
							},
							{
								Key:  "RENEWABLE",
								Path: "renewable",
								// Mode:
							},
						},
						// DefaultMode
					}
					if _, found := vaultSecret.Data[api.EnvVaultCACert]; found {
						volSrc.Items = append(volSrc.Items, v1.KeyToPath{
							Key:  api.EnvVaultCACert,
							Path: "ca.crt",
							// Mode:
						})
					}
					in.Spec.Template.Spec.Volumes = v1u.UpsertVolume(in.Spec.Template.Spec.Volumes, v1.Volume{
						Name: vaultSecret.Name,
						VolumeSource: v1.VolumeSource{
							Secret: &volSrc,
						},
					})
					for ci, c := range in.Spec.Template.Spec.Containers {
						c.Env = v1u.UpsertEnvVar(c.Env, v1.EnvVar{
							Name: api.EnvVaultAddress,
							ValueFrom: &v1.EnvVarSource{
								SecretKeyRef: &v1.SecretKeySelector{
									LocalObjectReference: v1.LocalObjectReference{
										Name: vaultSecret.Name,
									},
									Key: api.EnvVaultAddress,
								},
							},
						})
						if _, found := vaultSecret.Data[api.EnvVaultCACert]; found {
							c.Env = v1u.UpsertEnvVar(c.Env, v1.EnvVar{
								Name:  api.EnvVaultCAPath,
								Value: "/var/run/secrets/vaultproject.io/approle/ca.crt",
							})
						}
						in.Spec.Template.Spec.Containers[ci].Env = c.Env

						in.Spec.Template.Spec.Containers[ci].VolumeMounts = v1u.UpsertVolumeMount(c.VolumeMounts, v1.VolumeMount{
							Name:      vaultSecret.Name,
							MountPath: "/var/run/secrets/vaultproject.io/approle",
							ReadOnly:  true,
						})
					}

					return in
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}