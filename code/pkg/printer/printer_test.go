package printer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes/scheme"
)

func TestYamlPrinter(t *testing.T) {
	deploy := &appsv1.Deployment{}
	printer := printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.YAMLPrinter{})
	err := printer.PrintObj(deploy, os.Stdout)
	assert.NoError(t, err)
}

func TestJsonPrinter(t *testing.T) {
	deploy := &appsv1.Deployment{}
	printer := printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.JSONPrinter{})
	err := printer.PrintObj(deploy, os.Stdout)
	assert.NoError(t, err)
}

func TestTablePrinter(t *testing.T) {
	deploys := &metav1.Table{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ColumnDefinitions: []metav1.TableColumnDefinition{
			{
				Name: "Name",
				Type: "string",
			},
		},
		Rows: []metav1.TableRow{
			{
				Cells: []interface{}{"test"},
			},
		},
	}
	printer := printers.NewTypeSetter(scheme.Scheme).ToPrinter(printers.NewTablePrinter(printers.PrintOptions{}))
	err := printer.PrintObj(deploys, os.Stdout)
	assert.NoError(t, err)
}
