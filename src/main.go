package main

import (
	"fmt"
	"sync"
	"time"
)

type Cook struct {
	totalcook int
	speedcook int
}
type Customer struct {
	totaleat int
	speedeat int
}

const (
	m int = 3
	n int = 5
)

var (
	N         int = 30  //可存在的最大寿司量
	resource  int = 100 //材料
	sushi     chan int
	//tcook     [3]int = [3]int{700, 1000, 1300}
	//tcustomer        = [5]int{2000, 2000, 1600, 2000, 2500}
)
var wg sync.WaitGroup
var cook [m]Cook = [m]Cook{
	{30, 700},
	{40, 1000},
	{35, 1300},
}
var customer [n]Customer = [n]Customer{
	{15, 2000},
	{5, 2000},
	{25, 1600},
	{19, 2200},
	{30, 2500},
}

func main() {

	sushi = make(chan int, N)

	for i := 0; i < m; i++ {
		wg.Add(1)
		go makeSushi(i, sushi)
	}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go eatSushi(i, sushi)
	}
	wg.Wait()
}

func makeSushi(i int, sushi chan int) {
	defer wg.Done()
	for {
		if (resource > 0) && (cook[i].totalcook > 0) {
			sushi <- 1
			resource -= 1
			cook[i].totalcook -= 1
			switch i {
			case 0:
				fmt.Printf("cook %c makes a new sushi,\t total is %d \n", 'A', len(sushi))
				time.Sleep(time.Millisecond * 700)
			case 1:
				fmt.Printf("cook %c makes a new sushi,\t total is %d \n", 'B', len(sushi))
				time.Sleep(time.Millisecond * 1000)
			case 2:
				fmt.Printf("cook %c makes a new sushi,\t total is %d \n", 'C', len(sushi))
				time.Sleep(time.Millisecond * 1300)
			}
			//fmt.Printf("cook %d makes a new sushi,\t total is %d \n",i+1,len(sushi))
			//time.Sleep(time.Duration(cook[i].speedcook))
		} else {
			break
		}
	}
}

func eatSushi(i int, sushi chan int) {
	defer wg.Done()
	for {
		if customer[i].totaleat > 0 {
			_, ok := <-sushi
			customer[i].totaleat-=1
			if !ok {
				//close(sushi)
				break
			}
			switch i {
			case 0:
				fmt.Printf("customer %c has eated a sushi,left:%d\n", '1', len(sushi))
				time.Sleep(time.Millisecond * 1800)
			case 1:
				fmt.Printf("customer %c has eated a sushi,left:%d\n", '2', len(sushi))
				time.Sleep(time.Millisecond * 3000)
			case 2:
				fmt.Printf("customer %c has eated a sushi,left:%d\n", '3', len(sushi))
				time.Sleep(time.Millisecond * 1600)
			case 3:
				fmt.Printf("customer %c has eated a sushi,left:%d\n", '4', len(sushi))
				time.Sleep(time.Millisecond * 1700)
			case 4:
				fmt.Printf("customer %c has eated a sushi,left:%d\n", '5', len(sushi))
				time.Sleep(time.Millisecond * 2500)
			}
		}
	}
}
