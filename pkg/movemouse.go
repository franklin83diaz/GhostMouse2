package pkg

import (
	"time"

	"github.com/go-vgo/robotgo"
)

type mousePosition struct {
	X int
	Y int
}

var mousemovements []mousePosition = []mousePosition{}

// StartRecordingMouseMovement starts recording the mouse movements
func StartRecordingMouseMovement(stop chan bool) {
	mousemovements = []mousePosition{}
	for {

		x, y := robotgo.Location()
		mousemovements = append(mousemovements, mousePosition{X: x, Y: y})

		if <-stop {
			break
		}
	}
}

func PlayMouseMovement() {

	for _, movement := range mousemovements {
		time.Sleep(8 * time.Millisecond)
		go robotgo.Move(movement.X, movement.Y)
	}

}
