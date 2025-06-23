package repositories

type UserRepository interface {
	Create() (result string, err error)
	Update() (result string, err error)
	Delete() (result string, err error)
	Gets() (result string, err error)

	CreateAPIket() (result string, err error)
	UpdateAPIket() (result string, err error)
	DeleteAPIket() (result string, err error)
	GetsAPIket() (result string, err error)

	CreateMethod() (result string, err error)
	UpdateMethod() (result string, err error)
	DeleteMethod() (result string, err error)
	GetsMethod() (result string, err error)

	CreatePermission() (result string, err error)
	UpdatePermission() (result string, err error)
	DeletePermission() (result string, err error)
	GetsPermission() (result string, err error)
}
