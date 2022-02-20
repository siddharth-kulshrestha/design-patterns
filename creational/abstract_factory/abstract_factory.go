package main

import (
	"errors"
	"fmt"
)

// Abstract factory pattern lets you create a family of related objects
// Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.

// Lets say we have two factories, both produces shoe and shorts.
// adidas and nike

type iShoe interface {
	setSize(size int)
	setBrand(name string)
}

type iShort interface {
	setWaistSize(size int)
	setBrand(name string)
}

type iFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

type adidasShoe struct {
	size  int
	brand string
}

type nikeShoe struct {
	size  int
	brand string
}

func (a adidasShoe) setSize(size int) {
	a.size = size
}

func (a adidasShoe) setBrand(name string) {
	a.brand = name
}

func (n nikeShoe) setSize(size int) {
	n.size = size
}

func (n nikeShoe) setBrand(name string) {
	n.brand = name
}

type nikeShort struct {
	waistSize int
	brand     string
}

type adidasShort struct {
	waistSize int
	brand     string
}

func (n nikeShort) setWaistSize(size int) {
	n.waistSize = size
}

func (n nikeShort) setBrand(name string) {
	n.brand = name
}

func (a adidasShort) setWaistSize(size int) {
	a.waistSize = size
}

func (a adidasShort) setBrand(name string) {
	a.brand = name
}

type adidasFactory struct{}

func (adidasFactory) makeShoe() iShoe {
	return adidasShoe{
		brand: "adidas",
		size:  12,
	}
}

func (adidasFactory) makeShort() iShort {
	return adidasShort{
		waistSize: 45,
		brand:     "adidas",
	}
}

type nikeFactory struct{}

func (nikeFactory) makeShoe() iShoe {
	return nikeShoe{
		brand: "nike",
		size:  9,
	}
}

func (nikeFactory) makeShort() iShort {
	return nikeShort{
		brand:     "nike",
		waistSize: 36,
	}
}

func getSportsFactory(name string) (iFactory, error) {
	switch name {
	case "adidas":
		return &adidasFactory{}, nil
	case "nike":
		return &nikeFactory{}, nil
	}
	return nil, errors.New("name is invalid")
}

func main() {
	adidasFact, _ := getSportsFactory("adidas")
	nikeFact, _ := getSportsFactory("nike")
	adidasShoes := adidasFact.makeShoe()
	nikeShorts := nikeFact.makeShort()

	fmt.Println(adidasShoes)
	fmt.Println(nikeShorts)

}
