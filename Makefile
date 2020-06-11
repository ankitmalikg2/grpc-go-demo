gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean: 
	rm pb/*.go

server: 
	go run server.go

client :
	go run client.go