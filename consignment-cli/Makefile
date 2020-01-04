build:
#	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns \
		--network shippy_default \
		consignment-cli consignment.json \
		4e30c4a3-b280-4525-92b5-34c882d
