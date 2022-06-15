package pflag

import (
	"testing"

	"github.com/spf13/pflag"
)

var (
	Names []string
)

func TestPflags(t *testing.T) {
	flags := pflag.FlagSet{}
	flags.StringSliceVar(&Names, "names", []string{"*"}, "names joined with ,")

	args := []string{
		"--names=test1,test2",
	}
	flags.Parse(args)
	t.Log(Names)
}
