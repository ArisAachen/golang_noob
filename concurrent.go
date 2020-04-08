package main

import (
	"math/rand"
	"sync"
)

type nil_interface interface {
	Ping()
	Pang()
}

type nil_implement struct {
}

func (nil_inte nil_implement) Ping() {
	println("ping")
}

func (nil_inte *nil_implement) Pang() {
	println("pang")
}

var wg sync.WaitGroup

func main() {

	//空接口为nil的特殊情况
	var nil_parameter *nil_implement = nil

	var nil_inte nil_interface = nil_parameter

	if nil_inte == nil {
		println("nil")
	} else {
		println("not nil")
	}

	nil_inte.Pang()

	//关闭通道
	varChanInt := make(chan int, 15)
	varChanBool := make(chan bool)

	go writeChanTest(varChanInt)

	go closechantest(varChanInt, varChanBool)

	<-varChanBool
	println("chan end")

	//waitgroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go waitgrouptest()
	}

	wg.Wait()

	chanSelect := make(chan int)

	go selectTest(chanSelect)

	for i := 0; i < 10; i++ {
		println("select num = ", <-chanSelect)
	}

	//退出机制
	chanOut := make(chan int)
	chanRand := selectClose(chanOut)

	println(<-chanRand)

	close(chanOut)

	println(<-chanRand)

}

func writeChanTest(varChanInt chan int) {
	for i := 0; i < 10; i++ {
		varChanInt <- i
	}

	close(varChanInt)
	println("write end")
}

func closechantest(varChanInt chan int, varChanBool chan bool) {

	for {
		getInt, ok := <-varChanInt

		println("read_time = ", getInt)
		println("ok_mal = ", ok)

		if ok {
			println(getInt)
		} else {
			println("read end")
			varChanBool <- true
		}
	}
}

func waitgrouptest() {

	println("wait group")

	defer wg.Done()
}

func selectTest(chanInt chan int) {

	for {
		select {
		case chanInt <- 0:

			println("chanInt <- 0")
		case chanInt <- 1:
			println("chanInt <- 1")

		}
	}
}

func selectClose(chanInt chan int) chan int {

	ch := make(chan int)

	go func() {

	Label:
		for {

			select {
			case ch <- rand.Int():
				println("selectClose random")
			case <-chanInt:
				println("selectClose breakout")
				close(ch)
				break Label
			}

		}
	}()

	return ch
}
