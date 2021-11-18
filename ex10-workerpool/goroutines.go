package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"time"
)

func worker(id int, jobs <-chan float64) {
	fmt.Printf("worker:%d spawning\n", id)
	for j := range jobs {
		fmt.Printf("worker:%d sleep:%.1f\n", id, j)
		time.Sleep(time.Duration(float64(time.Second.Nanoseconds()) * j))
	}
	fmt.Printf("worker:%d stopping\n", id)
}

func Run(poolSize int) {
	jobs := make(chan float64, poolSize)
	id := 0

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i, _ := strconv.ParseFloat(s.Text(), 64)

		jobs <- i
		if id < poolSize {
			id++
			go worker(id, jobs)
		}
	}
	close(jobs)
}
