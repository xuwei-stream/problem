package question

import (
	"fmt"
	"testing"
	_ "testing"
)

func Test_rand13to5(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(rand13to5())
	}
}

func Test_rand5to13(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(rand5to13())
	}
}
