package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

/*
Select a random country, then select a random city.
*/
func random(cities map[string][]string) string {
	keys := reflect.ValueOf(cities).MapKeys()
	country := keys[rand.Intn(len(keys))].Interface()
	countryStr := fmt.Sprintf("%v", country)
	numCities := len(cities[countryStr])
	randomCityIndex := rand.Intn(numCities)
	return cities[countryStr][randomCityIndex]
}

//https://stackoverflow.com/questions/23482786/get-an-arbitrary-key-item-from-a-map
