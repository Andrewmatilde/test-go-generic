package main

import (
	"fmt"
	"testg/d"
	"testg/s"
)

func main() {
	nameClient := s.NameClient{}
	r := nameClient.Get(s.NameRequest{ID: 1})
	fmt.Println(r)

	nameClient2 := d.NameClient{
		DRequest: d.NewDRequest(d.NameRequest{ID: 1}, func() d.Name {
			return d.Name{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
			}
		}),
	}
	r2 := nameClient2.Get()
	fmt.Println(r2)
}
