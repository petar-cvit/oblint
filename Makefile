start:
	go run ./...

deps:
	docker-compose up -d

clean:
	docker rm $(docker ps -aq)
