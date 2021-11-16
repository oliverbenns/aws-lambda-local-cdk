build-lambda-image:
	docker build . -t=lambda-builder

build:
	go build -o bin/ping cmd/ping/main.go

up:
	docker compose up

deploy:
	cdk deploy
