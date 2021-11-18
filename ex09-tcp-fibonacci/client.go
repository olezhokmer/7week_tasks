package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {
	c, _ := net.Dial("tcp", "127.0.0.1:3000")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		json.NewEncoder(c).Encode(s.Text())
		answ := ""
		json.NewDecoder(c).Decode(&answ)

		fmt.Printf("%s", answ)
	}

	c.Close()
}
