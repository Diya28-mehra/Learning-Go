package main

import (
	"fmt"
	"sync"
)
func main(){
	//channels are a way through which multiple go routines can talk and communicate with each other 
	//may be they are waiting for some go routine they know in advance when one routine will end 
	fmt.Println("Channels in Go Lang")

	myCh:=make(chan int,1) //two parameters are provided to take more values for a single channel in buffer 
	wg:=&sync.WaitGroup{}
	// fmt.Println(<-myCh)
	// myCh<-5  // this will  give deadlock error

	wg.Add(2)
	go func(ch chan int,wg *sync.WaitGroup){
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh,wg)

	go func(ch chan int,wg *sync.WaitGroup){
		myCh<-5
		//myCh<-6 // too  many values will give error
		myCh<-6
		close(myCh)
		wg.Done()
	}(myCh,wg)

	wg.Wait()


}
