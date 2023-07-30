package constants

import (
	"errors"
)

var (
	ErrRouteNotFound = errors.New("1000")

	ErrDbError = errors.New("1001")
)

var ErrDescriptionByCode = map[error]string{
	ErrRouteNotFound: "Метод не найден",

	ErrDbError: "Ошибка при обращении к БД",
}
