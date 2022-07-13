package stackoverflow_go

// #cgo CFLAGS: -O0
/*
#include <stdio.h>
#include <stdint.h>

int recursive(uint64_t step) {
    step++;
	return recursive(step);
}
*/
import "C"

func recursiveCgo() {
	C.recursive(C.ulong(0))
}

func recursiveGo() {
	recursiveGo()
}
