build:
	docker build -t email-service .

run:
	docker run -p 50054:50051 \
		--network container-net \
		-e MICRO_SERVER_ADDRESS=:50054 \
		-e MICRO_REGISTRY=mdns \
		email-service
