package models

type OauthSource string

const (
	OauthSourceVk     OauthSource = "vk"
	OauthSourceMail   OauthSource = "mail"
	OauthSourceYandex OauthSource = "yandex"
)

type OauthUserInfo struct {
	ID    string
	Email string
}
