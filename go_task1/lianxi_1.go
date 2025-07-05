package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func main() {
	var array []uint = []uint{1, 1, 2, 4, 4, 5, 3, 3, 2, 6, 6}
	var array2 = []uint{1, 1, 2, 4, 5, 5, 3, 3, 2, 6, 6}
	array3 := []uint{1, 1, 5, 4, 4, 5, 3, 3, 2, 6, 6, 2}

	fmt.Println(queryOne(array))
	fmt.Println(queryOne(array2))
	fmt.Println(queryOne(array3))

	number := 12212
	var result bool = isPalindromeNumber(number)
	fmt.Println(result)

	printCount(100)

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
	fmt.Println(longestCommonPrefix([]string{"apple", "apple"}))           // "apple"
	fmt.Println(longestCommonPrefix([]string{"a"}))                        // "a"
	fmt.Println(longestCommonPrefix([]string{}))                           // ""
	fmt.Println(longestCommonPrefix([]string{"", "abc"}))
}

func queryOne(array []uint) uint {
	counts := make(map[uint]int, 15)
	for _, k := range array {
		counts[k]++
	}
	for k, _ := range counts {
		if counts[k] == 1 {
			return k
		}
	}
	return 0
}

// 回文数
func isPalindromeNumber(number int) bool {
	if number < 0 {
		return false
	}
	str := strconv.Itoa(number)
	bytes := []byte(str)
	for i := 0; i < len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-1-i] {
			return false
		}
	}
	return true
}

func isValid(s string) bool {
	if len(s)%2 > 0 {
		return false
	}
	paris := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if open, isFlag := paris[ch]; isFlag {
			if len(stack) == 0 || open != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

func swap(a, b int) (int, int) {
	swap := a
	a = b
	b = swap
	return a, b
}

func findMax(nums []int) int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// func countNum(str string) map[string]int {
// 	split := '-'
// 	bytes := []byte(str)
// 	m := make(map[string]int)
// 	for i:=0;i<len(bytes)-1;i++ {
// 		if bytes[i] == byte(split) {
// 			strTemp := bytes[]
// 			m[]
// 		}
// 	}
// }

type book struct {
	Title  string
	Author string
	Price  float64
}

func (b *book) info() {
	fmt.Print("《%s》 by %s：售价 %.2f", b.Title, b.Author, b.Price)
}

func printStr() {
	var wg sync.WaitGroup
	for i := 1; i < 6; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("I am goroutine #%d", num)
		}(i)
	}
	wg.Wait()
}

var count int
var lock sync.Mutex

func printCount(num int) {
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			count++
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	first := strs[0]
	if len(first) == 0 {
		return ""
	}
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || first[i] != strs[j][i] {
				return first[:i]
			}
		}
	}
	return first
}

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {

		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	result := make([]int, n+1)
	result[0] = 1
	return result
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i := 0
	for j := 1; j < n; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	res := [][]int{intervals[0]}

	for _, curr := range intervals {
		last := res[len(res)-1]
		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			res = append(res, curr)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}

		} else {
			m[num] = i
		}
	}
}
