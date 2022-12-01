package main

import (
	"fmt"
	"time"
)

func recoverFoo(i int) {
	recoverMessage := recover()
	if recoverMessage != nil {
		fmt.Println(recoverMessage, "-", i)
	}
	fmt.Printf("This is recoverFoo(%d)\n", i)
}

func foo(i int) {
	defer recoverFoo(i)
	panic("This is panic in foo()")
}

func main() {
	foo(0)

	fmt.Println("-----")

	for i := 1; i < 10; i++ {
		go foo(i)
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("This is finish main()")
}
