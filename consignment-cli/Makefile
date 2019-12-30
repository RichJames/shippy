build:
#	GOOS=linux GOARCH=amd64 go build
	docker build --no-cache -t consignment-cli .

run:
#	docker run -it -e MICRO_REGISTRY=mdns consignment-cli bash
	docker run -e MICRO_REGISTRY=mdns -it\
		--network container-net \
		consignment-cli sh
