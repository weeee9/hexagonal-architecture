package domain

type Customer struct {
	ID   int
	Name string
	City string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetByID(id int) (*Customer, error)
}

func NewCustomerRepository(driver ...string) CustomerRepository {
	var repo string

	if len(driver) > 0 {
		repo = driver[0]
	}

	switch repo {
	case "postgres":
		return NewCustomerRepotitoryPostgres()
	default:
		return NewCustomerRepotitoryStub()
	}
}
