package cgo

/*
int sum(int a, int b) {
  return a+b;
}
*/
import "C"

func sum(a, b int) int {
	return int(C.sum(C.int(a), C.int(b)))
}
