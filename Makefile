build-monitor:
	go build -o bin/outputs/default/main cmd/monitor/main.go

build-monitor-arm-v5:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/outputs/armv5/main cmd/monitor/main.go

run-monitor:
	go run cmd/monitor/main.go