.PHONY: build

build:
	docker build . -t thebillbot

run: 
	docker run --env-file ./prod.env thebillbot