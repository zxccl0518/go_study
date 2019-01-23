package main

import(
	"fmt";
	"time"
)

func main(){
	fmt.Println("hello")	

	// test_pipe()
	for i:= 1; i<100; i++{
		goroute_test(i)
	}

	time.Sleep(time.Second)
}


func test_pipe(){
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	fmt.Println(len(pipe))
	
	var t1 int 
	t1 =<- pipe

	var t2 int
	t2 =<- pipe

	fmt.Println( t1)
	fmt.Println( t2)
}

func goroute_test(a int){
	fmt.Println("a = ", a)
}