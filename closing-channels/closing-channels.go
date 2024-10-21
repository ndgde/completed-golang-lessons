package main

import (
	"fmt"
	"time"
)
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("received job", job)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	go func() {
		for job := 1; job <= 3; job++ {
			jobs <- job
			fmt.Println("sent job", job)
			time.Sleep(time.Second / 4)
		}
		close(jobs)
		fmt.Println("sent all jobs")
	}()

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
