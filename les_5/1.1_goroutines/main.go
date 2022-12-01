package main

import (
	"fmt"
	"runtime"
	"time"
)

func calc(num int) {
	fmt.Printf("Start func calc(%d)\n", num)

	var sum int
	for i := 0; i < num; i++ {
		sum += num + i
	}

	//Говорим планировщику Го, что нужно сменить горутину
	//Итог работы данной команды можно наблюдать в выводе терменала
	runtime.Gosched()

	fmt.Printf("Stop func calc(%d). Sum: %d\n", num, sum)
}

func main() {
	for i := 1; i < 6; i++ {
		go calc(i)
	}

	//Так как основной горутиной является main, то при ее завершении заканчивают работы остальные горутины.
	//Чтобы вызванные горутины успели закончить свое выполнение, мы можем прервать работу главной горутины.
	time.Sleep(100 * time.Millisecond)
}
