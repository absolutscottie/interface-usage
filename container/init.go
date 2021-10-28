package container

import (
	"github.com/absolutscottie/interface-usage/internal/dao"
	"github.com/absolutscottie/interface-usage/internal/postgres"
)

var (
	UserDAO dao.DAO
)

func Initialize() (err error) {
	UserDAO, err = postgres.NewDAO()
	if err != nil {
		return err
	}
	return nil
}
