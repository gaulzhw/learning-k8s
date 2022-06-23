package pflag

import (
	"testing"
	"time"

	"github.com/spf13/pflag"
)

var (
	names []string
	lasts time.Duration
)

func TestPflags(t *testing.T) {
	flags := pflag.FlagSet{}
	flags.StringSliceVar(&names, "names", []string{"*"}, "names joined with ,")
	flags.DurationVar(&lasts, "lasts", 0, "time duration")

	args := []string{
		"--names=test1,test2",
		"--lasts=3m",
	}
	flags.Parse(args)
	t.Log(names, lasts.Seconds())
}
