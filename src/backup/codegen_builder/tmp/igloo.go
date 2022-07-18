package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window" // replace this!!
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door" // replace this!!
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1 // replace this!!
}

func (b *iglooBuilder) getHouse() house {
	return house{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
