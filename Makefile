hello:
	echo "Hello"

build:
	go build -o bin/outputs/default/main main.go

build-arm-v5:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/outputs/armv5/main main.go

run:
	go run main.go