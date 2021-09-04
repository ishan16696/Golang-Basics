package main

import(
	"fmt"
	"time"
)

func help(){
	fmt.Printf("help called at %v\n",time.Now())
}


func main(){

	ticker := time.NewTicker(500*time.Millisecond)
	done := make(chan bool)

	go func(){

		for{
			select{
			case <-done:
				return
			case  <-ticker.C:
				//fmt.Printf("tick at %v\n",t)
				help()
			}
		}

	}()

	time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")

}