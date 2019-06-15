package hello

import ()


type user struct {
	ID   int64
	Name string
	Addr *address
}

type address struct {
	City   string
	ZIP    int
	LatLng [2]float64
}

var alex = user{}

// Hello writes a welcome string
func Hello() string {
	return "Hello, "+alex.Name
}
