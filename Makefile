
build:
	@go build -o ./bin/app.exe main.go

run: build
	@./bin/app.exe
