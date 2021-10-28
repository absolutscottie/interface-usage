package postgres

type DAOImpl struct {
	//implementation specific vars
}

func NewDAO() (*DAOImpl, error) {
	return &DAOImpl{}, nil
}

func (dao *DAOImpl) Create() error {
	return nil
}

func (dao *DAOImpl) Retrieve() error {
	return nil
}

func (dao *DAOImpl) Update() error {
	return nil
}

func (dao *DAOImpl) Delete() error {
	return nil
}
