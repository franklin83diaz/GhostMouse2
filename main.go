package main

import (
	"ghostmouse2/pkg"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	startText = "The record stops when\n you press click two times"
	stopText  = ""
)

func main() {

	recording := false
	text := stopText
	buttonText := "Start Recording"
	stop := make(chan bool)

	a := app.New()
	w := a.NewWindow("Ghost Mouse 2 - unpaid")

	w.Resize(fyne.NewSize(300, 100))
	var startRecordingButton *widget.Button
	startRecordingButton = widget.NewButton(buttonText, func() {

		if recording {
			recording = false
			startRecordingButton.SetText("Start Recording")
			stop <- true

		} else {
			recording = true

			startRecordingButton.SetText("Stop Recording")
			go pkg.StartRecordingMouseMovement(stop)
			go func() {
				for {
					stop <- false
					time.Sleep(8 * time.Millisecond)
				}
			}()
		}

	})

	replayButton := widget.NewButton("Replay", func() {
		println("Botón 2 presionado")
		pkg.PlayMouseMovement()
	})

	test := widget.NewLabel(text)

	buttonRow := container.NewHBox(startRecordingButton, replayButton, test)
	w.SetContent(buttonRow)

	w.ShowAndRun()
}
