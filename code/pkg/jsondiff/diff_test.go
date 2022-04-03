package jsondiff

import (
	"testing"

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

	diff(deploy, newDeploy)
}
