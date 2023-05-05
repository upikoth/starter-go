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
	sh scripts/migrateup.sh

migratedown:
	sh scripts/migratedown.sh

swagger:
	sh scripts/swagger.sh

lint:
	sh scripts/lint.sh
