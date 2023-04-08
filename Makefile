run:
	go run cmd/main.go

docker-rm:
	docker rm onelabhw2-app
	docker image rm onelabhw2-app

.PHONY: run docker-rm