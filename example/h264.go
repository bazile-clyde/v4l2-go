package main

import (
	"bytes"
	"fmt"
	"github.com/Charleye/v4l2-go"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	device := "/dev/video0"
	d, err := v4l2.Open(device)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	var cam v4l2.Camera
	cam.Device = *d

	cam.VerifyCaps()

	cam.Width = 800
	cam.Height = 600
	cam.PixelFormat = v4l2.GetFourCCByName("YUYV")
	cam.SetFormat()
	cam.AllocBuffers(4)
	cam.TurnOn()
	defer cam.TurnOff()

	f, _ := os.OpenFile("test_image.jpeg", os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()

	data := cam.Capture()
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err = jpeg.Encode(f, img, nil); err != nil {
		fmt.Println(err)
	}
}
