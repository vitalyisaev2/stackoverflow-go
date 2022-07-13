package stackoverflow_go

import (
	"fmt"
	"testing"
)

func TestStackOverflowGo(t *testing.T) {
	go func() {
		if err := trackStats("go"); err != nil {
			fmt.Println(err)
			t.Fatal(err)
		}
	}()

	recursiveGo()
}

func TestStackOverflowCgo(t *testing.T) {
	go func() {
		if err := trackStats("cgo"); err != nil {
			fmt.Println(err)
			t.Fatal(err)
		}
	}()

	recursiveCgo()
}
