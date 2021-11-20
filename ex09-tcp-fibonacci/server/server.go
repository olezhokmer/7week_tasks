package main

import (
	"encoding/json"
	"math/big"
	"net"

	"calc/models"
)

var cache = make(map[int]*big.Int)

func fibonacci(n int) *big.Int {
	if cache[n] != nil {
		return cache[n]
	}
	f2 := big.NewInt(0)
	f1 := big.NewInt(1)

	if n == 1 {
		return f2
	}

	if n == 2 {
		return f1
	}

	for i := 3; i <= n+1; i++ {
		next := big.NewInt(0)
		next.Add(f2, f1)
		f2 = f1
		f1 = next
	}
	cache[n] = f1
	return f1
}

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:3000")
	c, _ := l.Accept()
	var n int
	for {
		json.NewDecoder(c).Decode(&n)
		res := fibonacci(n)
		data := models.NewFib(res)
		json.NewEncoder(c).Encode(data)
	}
}
