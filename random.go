package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

/*
Select a random country, then select a random city.
*/
func random() string {
	//get country
	countryKeys := reflect.ValueOf(availableData.Cities).MapKeys()
	country := countryKeys[rand.Intn(len(countryKeys))].Interface()
	countryStr := fmt.Sprintf("%v", country)
	//get city
	numCities := len(availableData.Cities[countryStr])
	randomIndex := rand.Intn(numCities)
	return availableData.Cities[countryStr][randomIndex]
}
