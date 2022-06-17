package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const location = "/sys/class/leds/chromeos::kbd_backlight"

var brightnessLocation string = fmt.Sprintf("%s/%s", location, "brightness")
var maxBrightnessLocation string = fmt.Sprintf("%s/%s", location, "max_brightness")
var maxBrightness int
var brightness int

func main() {
	_brightness, err := readFileToInt(brightnessLocation)
	if err != nil {
		panic(err)
	}
	brightness = _brightness
	_maxBrightness, err := readFileToInt(maxBrightnessLocation)
	maxBrightness = _maxBrightness
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 1 {
		fmt.Print(brightness)
		os.Exit(0)
	}
	command := os.Args[1]
	amount := 5
	if len(os.Args) == 3 {
		amount, err = strconv.Atoi(os.Args[2])
		if err != nil {
			os.Exit(1)
		}
	}
	switch command {
	case "+":
		setBrightness(amount + brightness)
	case "-":
		setBrightness(amount - brightness)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}

func readFileToInt(path string) (int, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	fileString := strings.Trim(string(file), "\n")
	fileInt, err := strconv.Atoi(fileString)
	return int(fileInt), err
}

func setBrightness(brightness int) error {
	if brightness > maxBrightness {
		return nil
	}
	err := ioutil.WriteFile(maxBrightnessLocation, []byte(strconv.Itoa(brightness)), 0644)
	return err
}
