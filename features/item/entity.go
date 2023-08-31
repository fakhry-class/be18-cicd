package item

type Core struct {
	ID          uint
	Name        string
	UserID      uint
	Brand       string
	Description string
	Price       int
	Weight      int
	User        UserCore
}

type UserCore struct {
	ID    uint
	Name  string
	Email string
}

type ItemDataInterface interface {
	SelectAll(name string) ([]Core, error)
	// SelectByName(name string) ([]Core, error)
}

type ItemServiceInterface interface {
	GetAll(name string) ([]Core, error)
}
