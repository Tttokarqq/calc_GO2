package main

import(
	"sync"
	"fmt"
)

var(
	Buf []int
	mutex sync.Mutex
)

func Write(num int){
	mutex.Lock()
	defer mutex.Unlock()
	Buf = append(Buf, num)
}

func Consume() int{
	mutex.Lock()
	defer mutex.Unlock()
	a := buf[0]
	Buf = buf[1:]
	return a
}

func main(){
	Buf = []int{}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			Write(i)
		}()
	}
	wg.Wait()

	if len(Buf) != 1000 {
		fmt.Println(len(Buf))
	}
}