package jsondiff

import (
	"encoding/json"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestJson2Yaml(t *testing.T) {
	deploy := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
	}

	bytes, err := json.Marshal(deploy)
	assert.NoError(t, err)

	yamlBytes, err := yaml.JSONToYAML(bytes)
	assert.NoError(t, err)
	t.Log(string(yamlBytes))
}
