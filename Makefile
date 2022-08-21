BINARY_NAME = tic-tac-go
OUTPUT_FOLDER = bin/
OUTPUT_TARGET = ${OUTPUT_FOLDER}${BINARY_NAME}

build:
	@if [ ! -d ${OUTPUT_FOLDER} ] ; then mkdir -p ${OUTPUT_FOLDER} ; fi
	go build -o ${OUTPUT_TARGET} .


run:
	go run .


compile:
	@if [ ! -d ${OUTPUT_FOLDER} ] ; then mkdir -p ${OUTPUT_FOLDER} ; fi
	GOOS=linux GOARCH=386 go build -o ${OUTPUT_FOLDER}tic-tac-go-linux-i386 .
	GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_FOLDER}tic-tac-go-linux-amd64 .
	GOOS=linux GOARCH=arm go build -o ${OUTPUT_FOLDER}tic-tac-go-linux-arm .
	GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT_FOLDER}tic-tac-go-mac-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o ${OUTPUT_FOLDER}tic-tac-go-mac-arm64 .
	GOOS=windows GOARCH=386 go build -o ${OUTPUT_FOLDER}tic-tac-go-win-i386 .
	GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_FOLDER}tic-tac-go-win-amd64 .