package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	lobby_size = 3
)

type Customer struct {
	Id int
}

type Barber struct {
	free bool
}

func (c Customer) String() string {
	return fmt.Sprintf("%d", c.Id)
}

func startClients(customers chan *Customer) {
	count := 1
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		customers <- &Customer{count}
		count += 1
	}
}

func Haircut(barber *Barber, customer *Customer, finished chan *Barber) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Haircut done for client %s.\n", customer)
	finished <- barber
}

func startBarberShop(customers <-chan *Customer) {
	barber := &Barber{free: true}
	lobby := []*Customer{}
	syncBarberChan := make(chan *Barber)

	for {
		select {
		case customer := <-customers:
			if !barber.free {
				if len(lobby) < lobby_size {
					lobby = append(lobby, customer)
					fmt.Printf("customer %s sit on a free seat.\n", customer)
				} else {
					fmt.Printf("the client %s goes out.\n", customer)
				}
			} else {
				fmt.Printf("customer %s is going to barber.\n", customer)
				barber.free = false
				go Haircut(barber, customer, syncBarberChan)
			}
			fmt.Printf("lobby: %+v\n", lobby)
		case barber := <-syncBarberChan:
			if len(lobby) > 0 {
				customer := lobby[0]
				lobby = lobby[1:]
				fmt.Printf("client %s is going to barber.\n", customer)
				go Haircut(barber, customer, syncBarberChan)
			} else {
				barber.free = true
				fmt.Printf("barber is going to sleep.\n")
			}
			fmt.Printf("Lobby: %v\n", lobby)
		}
	}
}

func main() {
	customers := make(chan *Customer)
	go startClients(customers)
	go startBarberShop(customers)
	time.Sleep(1000 * time.Second)
}
