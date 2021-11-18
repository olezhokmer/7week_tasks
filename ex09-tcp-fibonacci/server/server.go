package main

import (
	"encoding/json"
	"math/big"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "127.0.0.1:3000")
	c, _ := ln.Accept()
	var str string
	for {
		json.NewDecoder(c).Decode(&str)
		n := new(big.Int)
		n, _ = n.SetString(str, 10)

		fib := fibonacci(n)
		json.NewEncoder(c).Encode("1ms " + fib.String() + "\n")
	}
}

func fibonacci(n *big.Int) *big.Int {
	f2 := big.NewInt(0)
	f1 := big.NewInt(1)

	if n.Cmp(big.NewInt(1)) == 0 {
		return f2
	}

	if n.Cmp(big.NewInt(2)) == 0 {
		return f1
	}

	for i := 3; n.Cmp(big.NewInt(int64(i))) >= 0; i++ {
		next := big.NewInt(0)
		next.Add(f2, f1)
		f2 = f1
		f1 = next
	}

	return f1
}
