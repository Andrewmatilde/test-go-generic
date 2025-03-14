package d

import "fmt"

type DReqInterface[Q, R any] interface {
	Get() R
}

type DRequest[Q, R any] struct {
	request   Q
	newObject func() R
}

func NewDRequest[Q, R any](request Q, newObject func() R) *DRequest[Q, R] {
	return &DRequest[Q, R]{
		request:   request,
		newObject: newObject,
	}
}

func (d *DRequest[Q, R]) Get() R {
	fmt.Println(d.request)
	return d.newObject()
}

type Name struct {
	ID        int
	FirstName string
	LastName  string
}

type NameRequest struct {
	ID int `json:"id"`
}

type NameClient struct {
	*DRequest[NameRequest, Name]
}
