package question

import (
	"math/rand"
)

/*
题目三
*/
//rand13() to rand5()，rand13()随机到0、11、12、13时重新rand，范围控制在1-10，1-5直接返回，6-10减去5返回
func rand13to5() int {
	var a int
	a = rand.Intn(13)
LOOP:
	for a > 10 || a == 0 {
		a = rand.Intn(13)
		goto LOOP
	}
	if a <= 5 {
		return a
	} else {
		return a - 5
	}
}

// rand5() to rand13()，两次rand5()，第一次只取[1,3]有效，1代表[1,5]区间，2代表[6-10]区间，3代表[11-15]区间，
//第二次rand5()表示在当前区间的第几个数，如果在[11,15]区间，只取[11,13]
func rand5to13() int {
	var first int
	first = rand.Intn(5)
LOOP:
	for first > 3 || first == 0 {
		first = rand.Intn(5)
		goto LOOP
	}

	var second int
	second = rand.Intn(5)
LOOP1:
	for second == 0 || (first == 3 && second > 3) {
		second = rand.Intn(5)
		goto LOOP1
	}
	return (first-1)*5 + second
}
