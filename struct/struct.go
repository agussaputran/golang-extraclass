package main

import "fmt"

// struct
type location struct {
	province string
	city     string
}

type contact struct {
	name    string
	email   string
	phone   string
	address location // embed struct
}

type person struct {
	name    string
	age     int
	address struct {
		street string
	}
}

func main() {
	var c1 contact
	c1.name = "agus"
	c1.email = "agus@gmail.com"
	c1.phone = "08882"
	c1.address.province = "DKI Jakarta"
	c1.address.city = "jakarta pusat"

	var c2 = contact{
		email: "saputra@gmail.com",
		name:  "saputra",
		phone: "9292827",
		address: location{
			city:     "Yogya",
			province: "DI Yogyakarta",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2.name)

	// anonymous struct
	var c3 = struct {
		name string
		age  int
	}{}

	c3.name = "agus"
	c3.age = 24

	//

	c4 := person{
		name: "imam",
		age:  22,
		address: struct{ street string }{
			street: "Jl. Durian",
		},
	}

	//
	// c4.name = "imam"
	// c4.age = 22
	// c4.address.street = "Jl. Mangga"

	fmt.Println(c4)

	var listContact = []contact{
		{name: "agus", email: "agus@gmail.com", phone: "272672",
			address: location{city: "jogja", province: "DIT"}},
		{name: "saputra", email: "saputra@gmail.com", phone: "272672",
			address: location{city: "jogja", province: "DIT"}},
		{name: "ntoi", email: "ntoi@gmail.com", phone: "272672",
			address: location{city: "jogja", province: "DIT"}},
	}

	listContact = append(listContact, contact{name: "fandy", email: "fandy@gmail.com", phone: "272672",
		address: location{city: "makassar", province: "Sulsel"}})

	fmt.Println("-------")
	fmt.Println(listContact)

	c1.Hello()

}

func (c contact) Hello() {
	fmt.Println("Hello ", c.name)
}
