package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func main() {

	var s1 []string

	for i := 1; i <= 10000000; i++ {
		s1 = append(s1, RandStringRunes(40))
	}

	//fmt.Println(s1)
	t1 := time.Now()
	sort.Strings(s1)
	t2 := time.Now()

	//fmt.Println(s1)

	fmt.Println("start: ", t1)
	fmt.Println("end: ", t2)
	fmt.Println(len(s1))

}
