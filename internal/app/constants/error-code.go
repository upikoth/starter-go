package constants

import (
	"errors"
)

var (
	ErrRouteNotFound     = errors.New("1000")
	ErrUserNotAuthorized = errors.New("1100")

	ErrUsersGetDbError = errors.New("1200")

	ErrUserGetIdNotExistInRequest = errors.New("1300")
	ErrUserGetIdNotNumber         = errors.New("1301")
	ErrUserGetNotFoundById        = errors.New("1302")
	ErrUserGetDbError             = errors.New("1303")
	ErrUserGetByEmailDbError      = errors.New("1304")
	ErrUserGetByEmailUserNotExist = errors.New("1305")

	ErrUserPostNoValidJson = errors.New("1400")
	ErrUserPostNoEmail     = errors.New("1401")
	ErrUserPostNoPassword  = errors.New("1402")
	ErrUserPostEmailExist  = errors.New("1403")
	ErrUserPostDbError     = errors.New("1404")
	ErrUserPostCreateHash  = errors.New("1405")

	ErrUserDeleteIdNotExistInRequest = errors.New("1500")
	ErrUserDeleteIdNotNumber         = errors.New("1501")
	ErrUserDeleteNotFoundById        = errors.New("1502")
	ErrUserDeleteDbError             = errors.New("1503")

	ErrUserPatchIdNotExistInRequest = errors.New("1600")
	ErrUserPatchIdNotNumber         = errors.New("1601")
	ErrUserPatchNoValidJson         = errors.New("1602")
	ErrUserPatchEmailExist          = errors.New("1603")
	ErrUserPatchDbError             = errors.New("1604")
	ErrUserPatchNotFoundById        = errors.New("1605")

	ErrSessionPostNoValidJson    = errors.New("1700")
	ErrSessionPostNoEmail        = errors.New("1701")
	ErrSessionPostNoPassword     = errors.New("1702")
	ErrSessionPostUserNotExist   = errors.New("1703")
	ErrSessionPostCreateJwtToken = errors.New("1704")
)

var ErrDescriptionByCode = map[error]string{
	ErrRouteNotFound:     "Метод не найден",
	ErrUserNotAuthorized: "Пользователь не авторизован",

	ErrUsersGetDbError: "Не удалось получить список пользователей",

	ErrUserGetIdNotExistInRequest: "Не передан id пользователя",
	ErrUserGetIdNotNumber:         "Переданный id должен быть числом",
	ErrUserGetNotFoundById:        "Пользователь с указанным id не найден",
	ErrUserGetDbError:             "Не удалось получить информацию о пользователе",
	ErrUserGetByEmailDbError:      "Не удалось получить информацию о пользователе",
	ErrUserGetByEmailUserNotExist: "Пользователя с такими email не существует",

	ErrUserPostNoValidJson: "Переданный json не валиден",
	ErrUserPostNoEmail:     "Не передан email пользователя",
	ErrUserPostNoPassword:  "Не передан пароль пользователя",
	ErrUserPostEmailExist:  "Пользователь с переданным email уже существует",
	ErrUserPostDbError:     "Не удалось создать пользователя",
	ErrUserPostCreateHash:  "Ошибка при создании пользователя",

	ErrUserDeleteIdNotExistInRequest: "Не передан id пользователя",
	ErrUserDeleteIdNotNumber:         "Переданный id должен быть числом",
	ErrUserDeleteNotFoundById:        "Пользователь с указанным id не найден",
	ErrUserDeleteDbError:             "Не удалось удалить пользователя",

	ErrUserPatchIdNotExistInRequest: "Не передан id пользователя",
	ErrUserPatchIdNotNumber:         "Переданный id должен быть числом",
	ErrUserPatchNoValidJson:         "Переданный json не валиден",
	ErrUserPatchEmailExist:          "Пользователь с переданным email уже существует",
	ErrUserPatchDbError:             "Не удалось обновить пользователя",
	ErrUserPatchNotFoundById:        "Пользователь с указанным id не найден",

	ErrSessionPostNoValidJson:    "Переданный json не валиден",
	ErrSessionPostNoEmail:        "Не передан email пользователя",
	ErrSessionPostNoPassword:     "Не передан пароль пользователя",
	ErrSessionPostUserNotExist:   "Пользователя с такими email и пароль не существует",
	ErrSessionPostCreateJwtToken: "Ошибка при входе в систему, попробуйте позже",
}
