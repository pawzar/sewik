package main

import (
	"fmt"
	"strconv"
	"time"

	"sewik/sync"
)

func digits(number int, dchnl chan int) {
	for number > 0 {
		dchnl <- number
		number /= 2
	}
	close(dchnl)
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}
func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main0() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

func main1() {
	n := 1024
	ch := make(chan int, 11)

	//go func() {
	for n > 0 {
		ch <- n
		n /= 2
	}
	close(ch)
	println("done")
	//}()

	for i := range ch {
		println(i)
	}
}
func main2() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func main() {
	wg := sync.LimitingWaitGroup{Limit: 5}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		println(i)
		go func(j int) {
			time.Sleep(time.Duration(j+1) * time.Second)
			println("< " + strconv.Itoa(j))
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		println(i)
		go func(j int) {
			time.Sleep(time.Duration(j+1) * time.Second)
			println("< " + strconv.Itoa(j))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
