.PHONY: build push all run

VERSION=1.0.0

all: build push
build:
	docker build . -t docker.io/kutrzeba/thebillbot:${VERSION}

run:
	docker run -d --restart on-failure --env-file prod.env kutrzeba/thebillbot:${VERSION}

push:
	docker push docker.io/kutrzeba/thebillbot:${VERSION}

pull:
	docker pull docker.io/kutrzeba/thebillbot:${VERSION}