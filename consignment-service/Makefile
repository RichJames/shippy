build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
#	docker build -t consignment-service .

runmongodb:
	docker run -d -p 27017:27017 mongo

run:
	docker run --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		consignment-service

createcon:
	curl -XPOST -H 'Content-Type: application/json' \
		-d '{"description": "This is a test", "weight": "500", "containers": []}' \
		http://localhost:8080/shippy/
