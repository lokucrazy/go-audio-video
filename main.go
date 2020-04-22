package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

func findBiggestDigit(num int, digit int) int {
	rem := num % digit
	if rem == num {
		return digit / 10
	}
	return findBiggestDigit(num, digit*10)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide audio file")
	}
	args := os.Args[1:]
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatalln("Could not open mp3 file:", err.Error())
	}
	defer file.Close()
	data, err := mp3.NewDecoder(file)
	if err != nil {
		log.Fatalln("Could not decode mp3:", err.Error())
	}
	audioBytes, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatalln("Could not read mp3:", err.Error())
	}
	fmt.Println(len(audioBytes))
	pixels := len(audioBytes) % findBiggestDigit(len(audioBytes), 1)
	upLeft := image.Point{0, 0}
	lowRight := image.Point{100, 100}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	index := 0
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if index < pixels {
				log.Println(index)
				img.Set(x, y, color.RGBA{audioBytes[index], audioBytes[index], audioBytes[index], 255})
				index++
			}
		}
	}

	file, err = os.Create("image.png")
	if err != nil {
		log.Fatalln("Could not create image:", err.Error())
	}
	defer file.Close()
	png.Encode(file, img)
}
