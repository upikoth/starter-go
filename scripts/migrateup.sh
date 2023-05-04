migrate -path migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_ADDR}/${DATABASE_NAME}?sslmode=disable" -verbose up
