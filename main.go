/*
You can add new devices by pasting that code into the end of file
for key, value := range sensors.Chips["Chip name"] {
			if key == "Chip key" {
				chip_name = value
			}
		}
*/

package main

import (
	"fmt"
	"github.com/ssimunic/gosensors"
	"time"
)

func main() {
	for true {
		sensors, err := gosensors.NewFromSystem()
		// sensors, err := gosensors.NewFromFile("/path/to/log.txt")

		if err != nil {
			panic(err)
		}
		var gpu string
		var cpu string
		var nvidia string
		for key, value := range sensors.Chips[""] { // Enter the name of pci device (you can see it by typing command 'sensors' in your terminal
			if key == "edge" {
				gpu = value //GPU dev
			}
		}
		for key, value := range sensors.Chips[""] {
			if key == "Tctl" {
				cpu = value //CPU dev
			}
		}
		for key, value := range sensors.Chips[""] {
			if key == "Composite" {
				nvidia = value //Discrete card dev
			}
		}
		// HERE you can paste new devices
		fmt.Printf("\rCPU: %s, Renoir: %s, Nvidia: %s", cpu, gpu, nvidia)
		time.Sleep(5 * time.Second) // You can change the value of sleeping
	}
}
