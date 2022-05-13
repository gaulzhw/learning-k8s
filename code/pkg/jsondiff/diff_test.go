package jsondiff

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wI2L/jsondiff"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestDiff(t *testing.T) {
	deploy := v1.Deployment{}
	newDeploy := deploy.DeepCopy()
	newDeploy.Spec.Template.Spec.InitContainers = []corev1.Container{
		{
			Name:    "init",
			Image:   "busybox",
			Command: []string{"/bin/sh", "-c", " echo 'init' && sleep 100 "},
		},
	}

	patch, err := jsondiff.Compare(deploy, newDeploy)
	assert.NoError(t, err)

	patchBytes, err := json.MarshalIndent(patch, "", "    ")
	assert.NoError(t, err)

	t.Logf("%s", patchBytes)
}
