client:
	go run client/main.go 

server:
	go run server/main.go 

client-tls:
	go run client/main.go 

server-tls:
	go run server/main.go 

cert:
	cd cert; ./gen.sh; cd ..