
build:
	@go build -o ./bin/app.exe main.go

run: build
	@./bin/app.exe

build:
	javac src/*.java -d ./bin

run: build
	java -cp ./bin $(program)	
