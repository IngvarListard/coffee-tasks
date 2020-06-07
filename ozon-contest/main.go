package main

import 	_ "net/http/pprof"

func resolve(r chan<- int, x int, f func(int) int) {
	r <- f(x)
}

func read(xStore chan<- int, in <-chan int, n int) {
	for i := 0; i < n; i++ {
		xStore <- <-in
	}
	close(xStore)
}


func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		var xx1 []int
		var xx2 []int
		x1ch := make(chan int, n)
		x2ch := make(chan int, n)
		go read(x1ch, in1, n)
		go read(x2ch, in2, n)

		rx1ch := make(chan int, n)
		rx2ch := make(chan int, n)
		for i := 0; i < n; i++ {
			x1 := <- x1ch
			x2 := <- x2ch
			xx1 = append(xx1, x1)
			xx2 = append(xx2, x2)
			go resolve(rx1ch, x1, f)
			go resolve(rx2ch, x2, f)
			res := <- rx1ch + <- rx2ch
			out <- res
		}
		close(rx1ch)
		close(rx2ch)
	}()
}

