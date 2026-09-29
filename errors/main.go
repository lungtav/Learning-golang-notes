package main

import (
	"errors"
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {

	totalCost := 0
	cost1, err := sendSMS(msgToCustomer)
	if err != nil {
		fmt.Println("message couldnt send", err)
		return 0, err
	}
	totalCost = totalCost + cost1

	cost2, err := sendSMS(msgToSpouse)
	if err != nil {
		fmt.Println("message couldnt send", err)
		return 0, err
	}
	totalCost = totalCost + cost2

	fmt.Println(totalCost)
	return totalCost, nil

}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

// custom errors
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

//test

func validateStatus(status string) error {
	if status == "" {
		return errors.New("ststus cannot be empty")
	}
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil

}

func main() {
	fmt.Println("Hello, world")
	sendSMSToCouple("Hello, how are you doing", "packs of water in fridge")

	//formatting strings
	const name = "Kim"
	const age = 22
	fmt.Printf("%v is %v years old.\n", name, age)

	//errors package
	err := errors.New("something went worng")
	fmt.Println(err)

	//test
	validateStatus("hello")
}
