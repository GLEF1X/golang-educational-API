package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (repo CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return repo.customers, nil
}

func NewCustomerRepositoryStub(initCustomers []Customer) *CustomerRepositoryStub {
	return &CustomerRepositoryStub{customers: initCustomers}
}
