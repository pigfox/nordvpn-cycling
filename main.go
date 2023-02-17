package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f: ", r)
		}
	}()
	for {
		cities := make(map[string][]string)
		countries := getCountries()
		for _, v := range countries {
			if contentMinLength < len(v) {
				cities[v] = getCities(v)
			}
		}

		connect(random(cities))
		connected := status()
		if connected {
			n := rand.Intn(maxSleepTimeMinutes)
			fmt.Println("Sleeping for " + strconv.Itoa(n) + " minutes...")
			time.Sleep(time.Duration(n) * time.Minute)
			clear()
		}
		fmt.Println("Reconnect attempt at: ", time.Now())
	}
}
