package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func SayHello(name string) (string, error) {

	//check if there is an Error

	if name == "" {
		return "", errors.New("Please provide a name")

	}

	randomMessage := RandomName()

	message := fmt.Sprintf(randomMessage, name)
	return message, nil

}

func RandomName() string {
	formates := []string{
		"hello %v your welcome",
		"hi Your most welcome,%v",
		"WhatUp %v your most welcome",
	}

	// this is equivelent of formates(3)
	return formates[rand.Intn(len(formates))]

}

func main() {

	name := "Kual"

	//calling the function by doing exeception handling
	message, err := SayHello(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)

}
