package once

import (
	"fmt"
	"testing"
)

func TestDone(t *testing.T) {
	o := Once{}
	t.Logf("once is done: %t", o.Done())

	f := func() error {
		fmt.Println("success")
		return nil
	}
	o.Do(f)
	t.Logf("once is done: %t", o.Done())
}
