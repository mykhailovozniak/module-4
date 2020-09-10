package module_4

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

func CreateInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Initializing...")
		instance = &Singleton{}
	})

	fmt.Println("Getting already initialized instance...")
	return instance
}

func SingletonClientCode()  {
	instance := CreateInstance()
	CreateInstance()
	CreateInstance()
	CreateInstance()
	CreateInstance()

	fmt.Println(instance)
}
