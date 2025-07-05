package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// i := 10
	// addTen(&i)
	// fmt.Println(i)

	// j := []int{1, 2, 5, 8}
	// mulTwo(&j)
	// for _, k := range j {
	// 	fmt.Println(k)
	// }
	// printNum()

	// rec := Rectangle{}
	// rec.Area()
	// rec.Perimeter()

	// ci := Circle{}
	// ci.Area()
	// ci.Perimeter()

	// emp := Employee{
	// 	EmployeeId: 10086,
	// 	P: Person{
	// 		Name: "张三",
	// 		Age:  18,
	// 	},
	// }

	// emp.PrintInfo()

	// channelOne()

	// channelTwo()

	cou := 7
	var wg sync.WaitGroup
	sync1(&cou, &wg)
	wg.Wait()
	fmt.Println(cou)

	sync2()
}

func addTen(i *int) {
	*i += 10
}

func mulTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func printNum() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		if i%2 == 1 {
			go func(num int) {
				defer wg.Done()
				fmt.Println(num)
			}(i)
		} else {
			go func(num int) {
				defer wg.Done()
				fmt.Println(num)
			}(i)
		}
	}
	wg.Wait()
}

func gorun(funs []func()) {
	var wg sync.WaitGroup
	for i, fun := range funs {
		wg.Add(1)
		go func(i int, task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Printf("任务 #%d 执行耗时: %v\n", i, duration)
		}(i, fun)
	}
	wg.Wait()
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (s *Rectangle) Area() {
	fmt.Println("Rectangle Area")
}

func (s *Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter")
}

type Circle struct {
}

func (s *Circle) Area() {
	fmt.Println("Circle Area")
}

func (s *Circle) Perimeter() {
	fmt.Println("Circle Perimeter")
}

type Employee struct {
	EmployeeId int
	P          Person
}

type Person struct {
	Name string
	Age  int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("名称：%s，年龄：%d，工号：%d", e.P.Name, e.P.Age, e.EmployeeId)
}

func channelOne() {
	ch := make(chan int)
	go func() {
		SendCh1(ch)
	}()

	ReadCh1(ch)
}

func SendCh1(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func ReadCh1(ch <-chan int) {
	for c := range ch {
		fmt.Println(c)
	}
}

func SendCh2(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func ReadCh2(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range ch {
		fmt.Println(c)
	}
}

func channelTwo() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(2)
	go SendCh2(ch, &wg)
	go ReadCh2(ch, &wg)
	wg.Wait()
}

var lock sync.Mutex

func sync1(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(count *int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lock.Lock()
				*count += 1
				lock.Unlock()
			}
		}(count)
	}
}

func sync2() {
	var count int64 = 5
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
