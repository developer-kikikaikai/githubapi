BIN=auth_server

all: build

build:
	GOPATH=~/go GO111MODULE=on go build -o ${BIN} -v main.go

clean:
	rm ${BIN}
