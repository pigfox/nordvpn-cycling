package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	setPotentialData()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f: ", r)
		}
	}()
	connectAttempts := 0
	for {
		if connectAttempts == config.MaxConnectAttempts {
			panic(strconv.Itoa(connectAttempts) + " retry attempts has reached max " + strconv.Itoa(config.MaxConnectAttempts))
		}
		getCountries()
		getCities()
		connect(random())
		connected := status()
		if connected {
			n := rand.Intn(config.MaxSleepTimeMinutes)
			fmt.Println("Sleeping for " + strconv.Itoa(n) + " minutes...")
			time.Sleep(time.Duration(n) * time.Minute)
			connectAttempts = 0
			clear()
		}
		connectAttempts++
		fmt.Println("Connect attempt #"+strconv.Itoa(connectAttempts)+" at ", time.Now())
	}
}
