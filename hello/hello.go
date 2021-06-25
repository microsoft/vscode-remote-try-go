package hello

// User user type
type User struct {
	ID   int64
	Name string
	Addr *Address
}

// Address address type
type Address struct {
	City   string
	ZIP    int
	LatLng [2]float64
}

var test = User{}

// Hello writes a welcome string
func Hello(name string) string {
	return "Hello, hello, " + name
}
