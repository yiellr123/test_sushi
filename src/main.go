package main

import (
	"fmt"
	"log"
	"os"
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

)
var wg sync.WaitGroup
var cook [m]Cook = [m]Cook{
	{30, 700},
	{40, 1000},
	{35, 1300},
}
var customer [n]Customer = [n]Customer{
	{15, 1800},
	{5, 3000},
	{25, 1600},
	{19, 1700},
	{30, 1500},
}

func main() {
	    file, err := os.OpenFile("a.log",  os.O_CREATE | os.O_WRONLY | os.O_APPEND,os.ModePerm)
	    if err != nil {
	        log.Fatalln(err)
	    }
	    logger := log.New(file, "", log.LstdFlags|log.Llongfile)

	sushi = make(chan int, N)

	for i := 0; i < m; i++ {
		wg.Add(1)
		go makeSushi(i, sushi,logger)
	}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go eatSushi(i, sushi,logger)
	}
	wg.Wait()
	fmt.Println("打烊")
	logger.Printf("打烊")
}

func makeSushi(i int, sushi chan int,logger *log.Logger) {
	defer wg.Done()
	for {
		if (resource > 0) && (cook[i].totalcook > 0) {
			sushi <- 1
			resource -= 1
			cook[i].totalcook -= 1
			fmt.Printf("cook %d makes a new sushi,\t total is %d \n",i+1, len(sushi))
			logger.Printf("cook %d makes a new sushi,\t total is %d \n", i+1, len(sushi))
			time.Sleep(time.Millisecond *time.Duration(cook[i].speedcook))

		} else {
			fmt.Println("有个厨子不干了")
			logger.Printf("有个厨子不干了")
			break
		}
	}
}

func eatSushi(i int, sushi chan int,logger *log.Logger) {
	defer wg.Done()
	for {
		if customer[i].totaleat > 0 {
			_, ok := <-sushi
			customer[i].totaleat-=1
			if !ok {
				//close(sushi)
				break
			}
			fmt.Printf("customer %d has eated a sushi,left:%d\n", i+1, len(sushi))
			logger.Printf("customer %d has eated a sushi,left:%d\n", i+1, len(sushi))
			time.Sleep(time.Millisecond * time.Duration(customer[i].speedeat))

		}else{
			break
		}
	}
	fmt.Println("有位顾客吃撑了")
	logger.Printf("有位顾客吃撑了")
}




//switch i {
//case 0:
//	fmt.Printf("cook %c makes a new sushi,\t total is %d \n", 'A', len(sushi))
//	logger.Printf("cook %c makes a new sushi,\t total is %d \n", 'A', len(sushi))
//	time.Sleep(time.Millisecond *time.Duration(cook[i].speedcook))
//case 1:
//	fmt.Printf("cook %c makes a new sushi,\t total is %d \n", 'B', len(sushi))
//	logger.Printf("cook %c makes a new sushi,\t total is %d \n", 'B', len(sushi))
//	time.Sleep(time.Millisecond * 1000)
//case 2:
//	fmt.Printf("cook %c makes a new sushi,\t total is %d \n", 'C', len(sushi))
//	logger.Printf("cook %c makes a new sushi,\t total is %d \n", 'C', len(sushi))
//	time.Sleep(time.Millisecond * 1300)
//}
//fmt.Printf("cook %d makes a new sushi,\t total is %d \n",i+1,len(sushi))
//time.Sleep(time.Duration(cook[i].speedcook))


//switch i {
//case 0:
//	fmt.Printf("customer %c has eated a sushi,left:%d\n", '1', len(sushi))
//	logger.Printf("customer %c has eated a sushi,left:%d\n", '1', len(sushi))
//	time.Sleep(time.Millisecond * 1800)
//case 1:
//	fmt.Printf("customer %c has eated a sushi,left:%d\n", '2', len(sushi))
//	logger.Printf("customer %c has eated a sushi,left:%d\n", '2', len(sushi))
//	time.Sleep(time.Millisecond * 3000)
//case 2:
//	fmt.Printf("customer %c has eated a sushi,left:%d\n", '3', len(sushi))
//	logger.Printf("customer %c has eated a sushi,left:%d\n", '3', len(sushi))
//	time.Sleep(time.Millisecond * 1600)
//case 3:
//	fmt.Printf("customer %c has eated a sushi,left:%d\n", '4', len(sushi))
//	logger.Printf("customer %c has eated a sushi,left:%d\n", '4', len(sushi))
//	time.Sleep(time.Millisecond * 1700)
//case 4:
//	fmt.Printf("customer %c has eated a sushi,left:%d\n", '5', len(sushi))
//	logger.Printf("customer %c has eated a sushi,left:%d\n", '5', len(sushi))
//	time.Sleep(time.Millisecond * 1500)
//}
