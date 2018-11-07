package server

import (
	"flag"
	"time"

	prom "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	"github.com/kubevault/operator/apis"
	cs "github.com/kubevault/operator/client/clientset/versioned"
	"github.com/kubevault/operator/pkg/controller"
	"github.com/kubevault/operator/pkg/docker"
	"github.com/spf13/pflag"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1"
)

type ExtraOptions struct {
	DockerRegistry     string
	MaxNumRequeues     int
	NumThreads         int
	QPS                float64
	Burst              int
	ResyncPeriod       time.Duration
	PrometheusCrdGroup string
	PrometheusCrdKinds prom.CrdKinds
}

func NewExtraOptions() *ExtraOptions {
	return &ExtraOptions{
		DockerRegistry:     docker.ACRegistry,
		MaxNumRequeues:     5,
		NumThreads:         2,
		QPS:                100,
		Burst:              100,
		ResyncPeriod:       10 * time.Minute,
		PrometheusCrdGroup: prom.Group,
		PrometheusCrdKinds: prom.DefaultCrdKinds,
	}
}

func (s *ExtraOptions) AddGoFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.DockerRegistry, "docker-registry", s.DockerRegistry, "Docker image registry for sidecar, init-container, check-job, recovery-job and kubectl-job")

	fs.Float64Var(&s.QPS, "qps", s.QPS, "The maximum QPS to the master from this client")
	fs.IntVar(&s.Burst, "burst", s.Burst, "The maximum burst for throttle")
	fs.DurationVar(&s.ResyncPeriod, "resync-period", s.ResyncPeriod, "If non-zero, will re-list this often. Otherwise, re-list will be delayed aslong as possible (until the upstream source closes the watch or times out.")

	fs.StringVar(&s.PrometheusCrdGroup, "prometheus-crd-apigroup", s.PrometheusCrdGroup, "prometheus CRD  API group name")
	fs.Var(&s.PrometheusCrdKinds, "prometheus-crd-kinds", " - EXPERIMENTAL (could be removed in future releases) - customize CRD kind names")

	fs.BoolVar(&apis.EnableStatusSubresource, "enable-status-subresource", apis.EnableStatusSubresource, "If true, uses sub resource for crds.")
}

func (s *ExtraOptions) AddFlags(fs *pflag.FlagSet) {
	pfs := flag.NewFlagSet("vault-server", flag.ExitOnError)
	s.AddGoFlags(pfs)
	fs.AddGoFlagSet(pfs)
}

func (s *ExtraOptions) ApplyTo(cfg *controller.Config) error {
	var err error

	cfg.DockerRegistry = s.DockerRegistry
	cfg.MaxNumRequeues = s.MaxNumRequeues
	cfg.NumThreads = s.NumThreads
	cfg.ResyncPeriod = s.ResyncPeriod

	cfg.ClientConfig.QPS = float32(s.QPS)
	cfg.ClientConfig.Burst = s.Burst

	if cfg.KubeClient, err = kubernetes.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.ExtClient, err = cs.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.CRDClient, err = crd_cs.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.PromClient, err = prom.NewForConfig(&s.PrometheusCrdKinds, s.PrometheusCrdGroup, cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.AppCatalogClient, err = appcat_cs.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	return nil
}
