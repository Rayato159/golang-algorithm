package main

import "fmt"

// Attribute
type House struct {
	Color    string
	Bedrooms uint8
	Toilets  uint8
}

type Condo struct {
	Color    string
	Bedrooms uint8
	Toilets  uint8
}

// Constructor
func NewHouse(color string, bedrooms uint8, toilets uint8) *House {
	return &House{
		Color:    color,
		Bedrooms: bedrooms,
		Toilets:  toilets,
	}
}

// Method
// Get && Set Color
func (h *House) SetColor(color string) {
	h.Color = color
}

func (h *House) GetColor() string {
	return h.Color
}

// Get && Set Bedrooms
func (h *House) SetBedrooms(bedrooms uint8) {
	h.Bedrooms = bedrooms
}

func (h *House) GetBedrooms() uint8 {
	return h.Bedrooms
}

// Get && Set Toilets
func (h *House) SetToilets(toilets uint8) {
	h.Toilets = toilets
}

func (h *House) GetToilets() uint8 {
	return h.Toilets
}

func main() {
	house := NewHouse("white", 2, 3)
	fmt.Printf("%v\n", house)

	house.SetToilets(1)
	fmt.Printf("%v\n", house)
	fmt.Println(house.GetColor())
}
