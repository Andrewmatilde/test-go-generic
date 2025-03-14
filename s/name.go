package s

type Name struct {
	ID        int
	FirstName string
	LastName  string
}

type NameRequest struct {
	ID int `json:"id"`
}

type NameInterface interface {
	Get(request NameRequest) Name
}

type NameClient struct{}

func (n *NameClient) Get(request NameRequest) Name {
	return Name{
		ID:        request.ID,
		FirstName: "John",
		LastName:  "Doe",
	}
}
