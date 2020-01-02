build:
	docker build -t user-client .

run:
	docker run -p 50051:50051 \
		--network container-net \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-client
