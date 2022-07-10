package question

import (
	"testing"
	_ "testing"
)

func TestPerm(t *testing.T) {
	main()
}

func TestPerm2(t *testing.T) {
	sum := 0
	Perm([]rune("ABC"), func(a []rune) {
		sum++
	})
	if sum != 6 {
		t.Fatalf("Perm method error, expect:%d,actual:%d", 6, sum)
	}
	t.Logf("test is success")
}
