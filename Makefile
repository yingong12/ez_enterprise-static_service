os?=linux
port?=8686


run: 
	go run main.go run -c .env -p $(port)

build:export GOOS=$(os)
build:export GOARCH=amd64
build:
	@echo "building binary for $(GOOS)..."
	go build -o ./static_service main.go
	@echo "done!"
	