package dao

type DAO interface {
	Create() error
	Retrieve() error
	Update() error
	Delete() error
}
