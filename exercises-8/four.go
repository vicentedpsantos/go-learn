package main

import (
	"fmt"
	"sort"
)

type ByInt []int

func (a ByInt) Len() int           { return len(a) }
func (a ByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInt) Less(i, j int) bool { return a[i] < a[j] }

type ByString []string

func (s ByString) Len() int           { return len(s) }
func (s ByString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByString) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Sort(ByInt(xi))
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Sort(ByString(xs))
	fmt.Println(xs)
}
