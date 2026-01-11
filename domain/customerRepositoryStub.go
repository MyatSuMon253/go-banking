package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:          "1",
			Name:        "John Doe",
			City:        "New York",
			Zipcode:     "10001",
			DateofBirth: "1990-01-01",
			Status:      "active",
		},
		{
			ID:          "2",
			Name:        "Jane Smith",
			City:        "Los Angeles",
			Zipcode:     "90001",
			DateofBirth: "1985-05-15",
			Status:      "inactive",
		},
	}

	return CustomerRepositoryStub{
		Customers: customers,
	}
}
