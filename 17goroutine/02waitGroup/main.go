package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 输出结果
	for a := 1; a <= 5; a++ {
		// 如果没有接受到值，就会造成main一直等待，直到完成接收到值
		fmt.Println(<-results)
	}
}
