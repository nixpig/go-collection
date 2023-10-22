package main

import (
	"errors"
	"fmt"

	"github.com/nixpig/go-collection/design-patterns/dependency-injection/data_store"
	"github.com/nixpig/go-collection/design-patterns/dependency-injection/logger"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) GetUserNameById(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

type SimpleLogic struct {
	l  logger.Logger
	ds data_store.DataStore
}

func (sl SimpleLogic) sayHello(userId string) (string, error) {
	name, ok := sl.ds.GetUserNameById(userId)
	if !ok {
		return "", errors.New("Could not find user")
	}

	greeting := "Hello, " + name + "!"

	sl.l.Log(greeting)

	return greeting, nil
}

func (sl SimpleLogic) sayGoodbye(userId string) (string, error) {
	name, ok := sl.ds.GetUserNameById(userId)
	if !ok {
		return "", errors.New("Could not find user")
	}

	greeting := "Goodbye, " + name + ""

	sl.l.Log(greeting)

	return greeting, nil
}

func NewSimpleLogic(l logger.Logger, ds data_store.DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func main() {
	simpleLogger := logger.Adapter(LogOutput)

	simpleDataStore := SimpleDataStore{
		userData: map[string]string{
			"1": "Bob",
			"2": "Mary",
			"3": "George",
			"4": "Andrew",
		},
	}

	logic := NewSimpleLogic(simpleLogger, simpleDataStore)

	logic.sayHello("2")   // => Hello, Mary!
	logic.sayGoodbye("3") // => Goodbye, George
}
