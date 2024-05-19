ifneq ("$(wildcard '.env')","")
include .env
endif

dev:
	sh scripts/dev.sh

build:
	sh scripts/build.sh

start:
	sh scripts/start.sh

generate:
	sh scripts/generate.sh

lint:
	sh scripts/lint.sh
