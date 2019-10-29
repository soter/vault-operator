/*
Copyright The KubeVault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubevault.dev/operator/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AWSAccessKeyRequests returns a AWSAccessKeyRequestInformer.
	AWSAccessKeyRequests() AWSAccessKeyRequestInformer
	// AWSRoles returns a AWSRoleInformer.
	AWSRoles() AWSRoleInformer
	// AzureAccessKeyRequests returns a AzureAccessKeyRequestInformer.
	AzureAccessKeyRequests() AzureAccessKeyRequestInformer
	// AzureRoles returns a AzureRoleInformer.
	AzureRoles() AzureRoleInformer
	// DatabaseAccessRequests returns a DatabaseAccessRequestInformer.
	DatabaseAccessRequests() DatabaseAccessRequestInformer
	// GCPAccessKeyRequests returns a GCPAccessKeyRequestInformer.
	GCPAccessKeyRequests() GCPAccessKeyRequestInformer
	// GCPRoles returns a GCPRoleInformer.
	GCPRoles() GCPRoleInformer
	// MongoDBRoles returns a MongoDBRoleInformer.
	MongoDBRoles() MongoDBRoleInformer
	// MySQLRoles returns a MySQLRoleInformer.
	MySQLRoles() MySQLRoleInformer
	// PostgresRoles returns a PostgresRoleInformer.
	PostgresRoles() PostgresRoleInformer
	// SecretEngines returns a SecretEngineInformer.
	SecretEngines() SecretEngineInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AWSAccessKeyRequests returns a AWSAccessKeyRequestInformer.
func (v *version) AWSAccessKeyRequests() AWSAccessKeyRequestInformer {
	return &aWSAccessKeyRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AWSRoles returns a AWSRoleInformer.
func (v *version) AWSRoles() AWSRoleInformer {
	return &aWSRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AzureAccessKeyRequests returns a AzureAccessKeyRequestInformer.
func (v *version) AzureAccessKeyRequests() AzureAccessKeyRequestInformer {
	return &azureAccessKeyRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AzureRoles returns a AzureRoleInformer.
func (v *version) AzureRoles() AzureRoleInformer {
	return &azureRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DatabaseAccessRequests returns a DatabaseAccessRequestInformer.
func (v *version) DatabaseAccessRequests() DatabaseAccessRequestInformer {
	return &databaseAccessRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GCPAccessKeyRequests returns a GCPAccessKeyRequestInformer.
func (v *version) GCPAccessKeyRequests() GCPAccessKeyRequestInformer {
	return &gCPAccessKeyRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GCPRoles returns a GCPRoleInformer.
func (v *version) GCPRoles() GCPRoleInformer {
	return &gCPRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MongoDBRoles returns a MongoDBRoleInformer.
func (v *version) MongoDBRoles() MongoDBRoleInformer {
	return &mongoDBRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MySQLRoles returns a MySQLRoleInformer.
func (v *version) MySQLRoles() MySQLRoleInformer {
	return &mySQLRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PostgresRoles returns a PostgresRoleInformer.
func (v *version) PostgresRoles() PostgresRoleInformer {
	return &postgresRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecretEngines returns a SecretEngineInformer.
func (v *version) SecretEngines() SecretEngineInformer {
	return &secretEngineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
