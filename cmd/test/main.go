package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func FilterOdd(nums []int) []int {
	//result := []int{}
	var result []int

	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	if len(result) == 0 {
		log.Println("here")
		return nil
	}
	return result
}

//func main() {
//	a := []int{1, 3, 5}
//	b := []int{2, 4, 6}
//	var c []int
//
//	fmt.Println("a:", FilterOdd(a)) // → [1 3 5]
//	fmt.Println("b:", FilterOdd(b)) // → nil
//	fmt.Println("c:", FilterOdd(c)) // → nil
//
//	r := FilterOdd(b)
//	if r == nil {
//		fmt.Println("Yes, result is nil!")
//	} else {
//		fmt.Println("Nope, result is not nil :(")
//	}
//}

//func main() {
//	defer func() {
//		r := recover()
//		log.Println("recover", r)
//	}()
//	defer fmt.Println("A")
//	os.Exit(1)
//	//panic("wewe")
//	log.Fatal("wewe")
//	fmt.Println("B")
//}

//func main() {
//	a := []int{1, 2, 3}
//	b := a[:2]
//	c := append(b, 99)
//
//	fmt.Println("a:", a)
//	fmt.Println("b:", b)
//	fmt.Println("c:", c)
//}

// 1, 2, 99
// 1, 2
// 1, 2, 99

//func main() {
//	var counter int
//
//	for i := 0; i < 5; i++ {
//		go func() {
//			for j := 0; j < 100; j++ {
//				counter++
//			}
//		}()
//	}
//
//	fmt.Println("Counter:", counter)
//}

func main() {

	runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(1 * time.Second)
}

//func main() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Recovered in f", r)
//		}
//	}()
//	defer fmt.Println("deferred 1")
//	fmt.Println("main")
//	log.Fatal("wewe")
//	defer fmt.Println("deferred 2")
//}
