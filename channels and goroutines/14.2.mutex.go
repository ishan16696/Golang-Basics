package main

import(
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func main(){
	rand.Seed(time.Now().UnixNano())

	count:=0
	finished:=0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		
		go func() {
			vote:= requestVote()

			mu.Lock()
			if vote{
				count++
			}
			finished++

			defer mu.Unlock()
		}()
	}

	for {

		mu.Lock()
		if count>=5 || finished==10{
			break
		}
		mu.Unlock()
	}

	if count>=5{
		fmt.Printf("Got majority..vote: %d\n",count)
	} else{
		fmt.Println("lost..")
	}


}


func requestVote() bool{
	time.Sleep(time.Duration(rand.Intn(100))* time.Millisecond)
	return rand.Int() %2 ==0
}