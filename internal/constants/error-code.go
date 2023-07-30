package constants

import (
	"errors"
)

var (
	ErrRouteNotFound = errors.New("1000")

	ErrDbError    = errors.New("1001")
	ErrDbNotFound = errors.New("1002")

	ErrRequestValidation = errors.New("1003")
)

var ErrDescriptionByCode = map[error]string{
	ErrRouteNotFound: "Метод не найден",

	ErrDbError:    "Ошибка при обращении к БД",
	ErrDbNotFound: "Не найдено в БД",

	ErrRequestValidation: "Переданы некорректные параметры запроса",
}
