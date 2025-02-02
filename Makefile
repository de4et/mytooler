path-copy: 
	copy /Y .\bin\app.exe "%TOOLS%\mytooler.exe"

# go --------------------
build:
	@go build -o ./bin/app.exe main.go

run: build
	@./bin/app.exe
