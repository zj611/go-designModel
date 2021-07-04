package testing

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var once sync.Once

type single struct {
	Name string
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating a singleInstance")
				singleInstance = &single{"singleInstance"}
			})
	} else {
		fmt.Println("SingleInstance already created.")
	}
	return singleInstance
}

func TestSingleInstance(t *testing.T) {
	for i := 0; i < 5; i++ {
		go getInstance()
	}
	u := getInstance()
	fmt.Println("name:", u.Name)
	time.Sleep(1 * time.Minute)

}
