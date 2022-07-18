package main

import "fmt"

func main() {

	normalBuilder := getBuilder("normal")
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	fmt.Printf("Normal House WindowType: %v\n", normalHouse.windowType)
	fmt.Printf("Normal House DoorType: %v\n", normalHouse.doorType)
	fmt.Printf("Normal House NumFloor: %v\n", normalHouse.floor)

	iglooBuilder := getBuilder("igloo")
	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Printf("Igloo House WindowType: %v\n", iglooHouse.windowType)
	fmt.Printf("Igloo House DoorType: %v\n", iglooHouse.doorType)
	fmt.Printf("Igloo House NumFloor: %v\n", iglooHouse.floor)

}
