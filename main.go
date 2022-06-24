package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	fmt.Println("xD")

	donationUrl := "https://widgets.tipply.pl/LATEST_MESSAGES/a067eac6-4f9d-4cd7-ad7a-955672416d33/e743c6dd-ce3d-4f59-b1bf-360576bfc939"
	selector := "div.sc-gqjmRU:nth-child(1)"

	rpi := raspi.NewAdaptor()
	lights := gpio.NewDirectPinDriver(rpi, "2")

	lights.On()
	time.Sleep(time.Second * 3)
	lights.Off()

	page := rod.New().MustConnect().MustPage(donationUrl)
	donation := ""
	for {
		el, err := page.Element(selector)

		if err != nil {
			fmt.Println("no el")
		}

		text, err := el.Text()
		if err != nil {
			fmt.Println("no text")
		}
		// fmt.Println(text, donation)
		if text != donation {
			text = donation

			lights.On()
			time.Sleep(time.Second * 5)
			lights.Off()

		}

		time.Sleep(time.Second)
	}

}
