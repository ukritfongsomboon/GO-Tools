package repositories

type MenuRepository interface {
	Create() (result string, err error)
	Update() (result string, err error)
	Delete() (result string, err error)
	Gets() (result string, err error)
}
