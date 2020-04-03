package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"sewik/pkg/sync"
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
	ch := make(chan string, 9)
	ch <- "naveen"
	fmt.Printf("%d/%d\n", cap(ch), len(ch))
	ch <- "paul"
	fmt.Printf("%d/%d\n", cap(ch), len(ch))
	ch <- "steve"
	fmt.Printf("%d/%d\n", cap(ch), len(ch))
	ch <- "alan"
	fmt.Printf("%d/%d\n", cap(ch), len(ch))
}
func main3() {
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
func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bc := widgets.NewBarChart()
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
