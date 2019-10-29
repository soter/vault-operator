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
package google

import (
	"testing"

	api "kubevault.dev/operator/apis/kubevault/v1alpha1"
	"kubevault.dev/operator/pkg/vault/util"

	"github.com/stretchr/testify/assert"
	core "k8s.io/api/core/v1"
)

func TestOptions_Apply(t *testing.T) {
	expected := []string{
		"--mode=google-cloud-kms-gcs",
		"--google.storage-bucket=test",
		"--google.kms-project=pro",
		"--google.kms-location=loc",
		"--google.kms-key-ring=r-key",
		"--google.kms-crypto-key=c-key",
	}
	cont := core.Container{
		Name: util.VaultUnsealerContainerName,
	}
	pt := &core.PodTemplateSpec{
		Spec: core.PodSpec{
			Containers: []core.Container{cont},
		},
	}

	opts, err := NewOptions(api.GoogleKmsGcsSpec{
		KmsCryptoKey: "c-key",
		KmsKeyRing:   "r-key",
		KmsLocation:  "loc",
		KmsProject:   "pro",
		Bucket:       "test",
	})
	assert.Nil(t, err)

	err = opts.Apply(pt)
	assert.Nil(t, err)

	assert.Equal(t, expected, pt.Spec.Containers[0].Args)
}
