package util

import (
	"log"
	"testing"
	"time"
)

func TestGoFunc(t *testing.T) {
	start := time.Now()
	for i := 0; i < 20; i++ {
		arr := []int{1, 2, 3}
		GoFn("worker", 50, true, worker, arr, "1.01", "abc")
	}
	for i := 0; i < 30; i++ {
		arr := []int{1, 2, 3}
		GoFn("worker1", 10, true, worker, arr, "2.02", "def")
	}
	log.Println("complete", time.Since(start).Seconds())
	time.Sleep(10 * time.Second)
}

func worker(done <-chan int, id int, args ...interface{}) {
	p0 := args[0].([]int)[1]
	p1, p2 := args[1], args[2]
	fPln("id:", id, " --- ", p0, p1, p2)
	// fPln(p0)
	// fPln(p1)
	// fPln(p2)

	time.Sleep(2 * time.Second)
	<-done
}
