package ex01_test

import (
	"example.com/ex01"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func ExampleDescribePlant_UnknownPlant() {
	plant := UnknownPlant{
		FlowerType: "Rose",
		LeafType:   "Lanceolate",
		Color:      10,
	}
	ex01.DescribePlant(plant)
	// Output:
	// FlowerType:Rose
	// LeafType:Lanceolate
	// Color(color_scheme=rgb):10
}

func ExampleDescribePlant_AnotherUnknownPlant() {
	plant := AnotherUnknownPlant{
		FlowerColor: 20,
		LeafType:    "Oval",
		Height:      15,
	}
	ex01.DescribePlant(plant)
	// Output:
	// FlowerColor:20
	// LeafType:Oval
	// Height(unit=inches):15
}

func ExampleDescribePlant_EmptyStruct() {
	type EmptyPlant struct{}
	plant := EmptyPlant{}
	ex01.DescribePlant(plant)
	// Output:
	// Empty struct
}

func ExampleDescribePlant_NotStruct() {
	plant := 1
	ex01.DescribePlant(plant)
	// Output:
	// Expected struct
}
