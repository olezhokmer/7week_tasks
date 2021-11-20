package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"calc/models"
)

func main() {
	c, _ := net.Dial("tcp", "127.0.0.1:3000")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		start := time.Now()
		n, _ := strconv.ParseInt(s.Text(), 10, 64)
		json.NewEncoder(c).Encode(n)
		var answ models.Fib
		json.NewDecoder(c).Decode(&answ)
		fmt.Printf("%s %d\n", time.Since(start), answ.Val)
	}

	c.Close()
}
