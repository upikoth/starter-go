package constants

import (
	"errors"
)

var (
	ErrRouteNotFound = errors.New("1000")

	ErrUsersGetDbError = errors.New("1100")
)

var ErrDescriptionByCode = map[error]string{
	ErrRouteNotFound: "Метод не найден",
}
