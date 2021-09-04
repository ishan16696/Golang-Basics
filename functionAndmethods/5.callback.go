package main


import(
	"fmt"
)

type LeaderCallbacks struct {
	OnStartedLeading func(s string)	
	OnStoppedLeading func()
}

type LeaderElector struct {
	Callbacks *LeaderCallbacks
	name string
}


func backupSidecar(){

	leaderCallbacks := &LeaderCallbacks{
		OnStartedLeading: func(s string) {
			fmt.Println("Started leading...",s)
			
		},
		OnStoppedLeading: func() {
			fmt.Println("Stopped leading...")	
		},
	}

	a:=LeaderElector{
		Callbacks: leaderCallbacks,
		name: "ISHAN",
	}

	a.Callbacks.OnStartedLeading("ISHAN")
	a.Callbacks.OnStoppedLeading()
	
}

func main(){
	fmt.Println("main Started...")
	backupSidecar()
	fmt.Println("main stopped..")

}

