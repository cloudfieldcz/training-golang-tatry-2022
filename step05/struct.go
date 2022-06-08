package main

type CustomStruct struct {
	Name string
}

// Validate does nothing for this example.
func (c CustomStruct) String() string {
	return "Name: " + c.Name
}

type AbstractStruct interface {
	String() string
}
