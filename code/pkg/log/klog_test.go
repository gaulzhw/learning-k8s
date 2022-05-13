package log

import (
	"flag"
	"testing"

	"k8s.io/klog/v2"
)

func TestKLog(t *testing.T) {
	klog.InitFlags(nil)
	flag.Parse()
	defer klog.Flush()

	klog.Info("just for test")
}
