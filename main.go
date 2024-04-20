package main

import "fmt"

func main() {
	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- "Every 500ms"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		c2 <- "Every 2s"
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()

	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		println(msg1)
	// 	case msg2 := <-c2:
	// 		println(msg2)
	// 	}
	// }

	jobs := make(chan int, 100)
	result := make(chan int, 100)

	go worker(jobs, result)
	go worker(jobs, result)
	go worker(jobs, result)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)
	for i := 0; i < 100; i++ {
		fmt.Println(<-result)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
