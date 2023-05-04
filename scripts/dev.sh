# nodemon обеспечивает live reload сервера
# для запуска скрипта он должен быть установлен глобально
# npm install -g nodemon
nodemon --exec go run cmd/apiserver/main.go --signal SIGTERM
