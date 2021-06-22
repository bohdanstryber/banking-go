package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"101", "John", "New York", "10293", "1990-01-29", "1"},
		{"102", "Smith", "Alaska", "22144", "1995-05-29", "1"},
		{"101", "Jenifer", "New York", "10293", "1990-01-29", "1"},
	}

	return CustomerRepositoryStub{customers}
}
