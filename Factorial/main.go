package main
import (
	"Factorial/Client"
	"Factorial/server"
	"time"
)

func main(){
	go func() {
		server.Start("50063")
	}()
    time.Sleep(2*time.Second)
	Client.Start("50063")
}
