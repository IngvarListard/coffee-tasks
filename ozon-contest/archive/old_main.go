package main

func executeWorker(f func(int) int, x1 int, res chan<- int) {
	res <- f(x1)
}

// test -> github.com/qypec/coffee-tasks/tree/master/merge_channels
func merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			res := make(chan int, 2)
			go executeWorker(f, x1, res)
			go executeWorker(f, x2, res)

			sum := 0
			for j := 0; j < 2; j++ {
				sum += <-res
			}
			out <- sum
			close(res)
		}
	}()
}

