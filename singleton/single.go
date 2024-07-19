package singleton

/*
	create Only a instance for any
*/

import (
	"fmt"
	"sync"
)

var Lock = &sync.Mutex{}

type Single struct {
}

var SingleInstance *Single

func GetInstance() *Single {
	if SingleInstance == nil {
		Lock.Lock()
		defer Lock.Unlock()
		if SingleInstance == nil {
			fmt.Println("Creating single instance now.")
			SingleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return SingleInstance
}
