build:
	protoc -I. \
		--go_out=plugins=micro:. \
		proto/auth/auth.proto

	docker build -t shippy-user-service .

runpostgres:
	docker run -d \
		-p 5432:5432 \
		postgres

run:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASSWORD=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		shippy-user-service

runold:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_NAME=userdb \
		-e DB_PASSWORD=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		shippy-user-service

callapi:
	curl -XPOST -H 'Content-Type: application/json' \
		-d '{"id":"b5652d08-b628-4158-8039-566cb7156b41"}' \
		http://localhost:8080/auth/Get

callrpc:
	curl -XPOST -H 'Content-Type: application/json' \
		-d '{"service": "shippy.auth", "method": "Auth.Auth", "request": {"email":"fff@gmail.com", "password", "yabadabado"}}' \
		http://localhost:8080/rpc
