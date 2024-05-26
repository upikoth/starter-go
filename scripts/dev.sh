# nodemon обеспечивает live reload сервера
# для запуска скрипта он должен быть установлен глобально
# npm install -g nodemon
nodemon --exec go run cmd/app/main.go --signal SIGTERM --quiet  --ignore authorized_key.json
