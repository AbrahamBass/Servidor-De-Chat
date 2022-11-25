package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct{}

var db *DataBase
var lock sync.Mutex

func (DataBase) CreatingSingleConnection() {
	fmt.Println("Creating Singleton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creating Done")
}

func getDataBaseIntace() *DataBase {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating db connection")
		db = &DataBase{}
		db.CreatingSingleConnection()
	} else {
		fmt.Println("db Alredy Created")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func () {
			defer wg.Done()
			getDataBaseIntace()
		}()
	}

	wg.Wait()
}
