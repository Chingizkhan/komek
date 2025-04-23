package main

import (
	"fmt"
	"log"
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

//func main() {
//
//	runtime.GOMAXPROCS(1)
//	for i := 0; i < 5; i++ {
//		i := i
//		go func() {
//			fmt.Println(i)
//		}()
//	}
//
//	time.Sleep(1 * time.Second)
//}

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

//func main() {
//	arr := []int{1, 2, 3}
//
//	operation1(arr[0:2])
//	fmt.Println(arr)
//
//	operation2(arr[0:2]) // len = 2, cap =3
//	fmt.Println(arr)
//}
//
//func operation1(arr []int) {
//	arr = append(arr, 4)
//}
//
//func operation2(arr []int) {
//	arr = append(arr, 5, 6)
//}

//func HasDuplicates(items []int) bool {
//	// true — если есть дубликаты
//}

//func main() {
//	fmt.Println(sum(5, 6))
//}
//
//func sum(a, b int) (c int) {
//	defer func() {
//		c++
//	}()
//	c = a + b
//
//	return c
//}

//func main() {
//	arr := []int{1, 2, 3} // len=3, cap=3
//
//	operation1(arr[0:2]) // len=2, cap=3
//	fmt.Println(arr)     // [1, 2, 4]
//
//	operation2(arr[0:2])
//	fmt.Println(arr)
//}
//
//func operation1(arr []int) {
//	arr = append(arr, 4)
//}
//
//func operation2(arr []int) {
//	arr = append(arr, 5, 6)
//}

func CharFrequency(s string) map[rune]int {
	result := make(map[rune]int, len(s))

	// TODO: реализуй подсчёт символов
	for _, c := range s {
		result[c] = result[c] + 1
	}

	return result
}

func main() {
	text := "banana"
	freq := CharFrequency(text)
	fmt.Println(freq)
}
