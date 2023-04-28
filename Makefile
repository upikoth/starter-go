ifneq ("$(wildcard '.env')","")
include .env
endif

dev:
	sh scripts/dev.sh

build:
	sh scripts/build.sh

start:
	sh scripts/start.sh

migrateup:
	migrate -path migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_ADDR}/${DATABASE_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_ADDR}/${DATABASE_NAME}?sslmode=disable" -verbose down -all
