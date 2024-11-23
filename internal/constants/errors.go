package constants

import "errors"

var (
	ErrDBEntityNotFound = errors.New("сущность не найдена")

	ErrSessionNotFound                 = errors.New("сессия не найдена")
	ErrSessionCreateInvalidCredentials = errors.New("некорректные данные для создания сессии")

	ErrUserNotFound     = errors.New("пользователь не найден")
	ErrUserAlreadyExist = errors.New("пользователь уже существует")

	ErrRegistrationNotFound = errors.New("регистрация не найдена")

	ErrPasswordRecoveryRequestNotFound = errors.New("запрос на восстановление пароля не найден")
)
