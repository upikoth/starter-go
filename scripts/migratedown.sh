if test -f ".env"; then
	set -o allexport
	source .env
	set +o allexport
fi

migrate -path migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_ADDR}/${DATABASE_NAME}?sslmode=disable" -verbose down 1
