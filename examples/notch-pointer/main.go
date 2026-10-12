// Notch-pointer prints the pointer's position as it moves anywhere on the
// screen, while another app is in front too, and quits after 10 seconds.
//
//	go run ./examples/notch-pointer
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/egoist/mygo"
)

func main() {
	mygo.App.SetActivationPolicy(mygo.ActivationPolicyAccessory)
	mygo.App.WhenReady(func() {
		off := mygo.App.OnGlobalPointer(func(x, y int) { fmt.Println("pointer:", x, y) })
		time.AfterFunc(10*time.Second, func() {
			off()
			mygo.App.Quit()
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
