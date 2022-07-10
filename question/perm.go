package question

import "fmt"

/*
题目一
*/
// Perm() 对 a 形成的每一排列调用 f().
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
func perm(a []rune, f func([]rune), i int) {
	if i == len(a) {
		f(a)
	}
	for j := i; j < len(a); j++ {
		swap(a, i, j)
		perm(a, f, i+1)
		swap(a, i, j)
	}
}

func swap(a []rune, i int, j int) {
	var tmp rune
	tmp = a[i]
	a[i] = a[j]
	a[j] = tmp
}

func main() {
	Perm([]rune("ABC"), func(a []rune) {
		fmt.Println(string(a))
	})
}
