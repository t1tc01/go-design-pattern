package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstancev2() *single {
	if singleInstance == nil {
		// đảm bảo rằng khối mã khởi tạo instance duy nhất (singleInstance)
		// chỉ được thực thi một lần, ngay cả khi nhiều goroutine gọi hàm GetInstancev2 đồng thời
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
