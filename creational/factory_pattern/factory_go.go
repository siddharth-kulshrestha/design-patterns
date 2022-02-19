package main

import "errors"

// Factory pattern is a creational because it provides a way to create objects without exposing the
// creational logic to the client.

// In OOPs: Factory Method defines a method, which should be used for creating objects instead of direct constructor call (new operator).
// Subclasses can override this method to change the class of objects that will be created.

// It is impossible to implement classic factory pattern in golang due to lack of classic OOPS features like class and inheritance.
// However, we can still implement the basic version of the pattern, the Simple Factory.

// factory defines an interface which will serve as an abstraction over all objects created via factory
type factory interface {
	setA(a string)
	setB(b string)
	getA() string
	getB() string
}

// This is factoryImpl
type factoryImpl struct {
	A string
	B string
}

func (f factoryImpl) setA(a string) {
	f.A = a
}

func (f factoryImpl) setB(b string) {
	f.B = b
}

func (f factoryImpl) getA() string {
	return f.A
}

func (f factoryImpl) getB() string {
	return f.B
}

type factoryImplObjOne struct {
	factoryImpl
}

func newFactoryImplObjOne() factory {
	return factoryImpl{
		A: "Obj One",
		B: "Obj 2",
	}
}

type factoryImplObjTwo struct {
	factoryImpl
}

func newFactoryImplObjTwo() factory {
	return factoryImpl{
		A: "Obj Two",
		B: "Obj 2",
	}
}

func getFactoryObj(idx int) (factory, error) {

	switch idx {
	case 1:
		return newFactoryImplObjOne(), nil
	case 2:
		return newFactoryImplObjTwo(), nil
	}
	return nil, errors.New("idx doesn't matches with any implementation of factory")

}
