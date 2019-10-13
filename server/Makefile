BIN=auth_server

all: build

build:
	GOPATH=~/go GO111MODULE=on ${GOPATH}/bin/go1.13 build -o ${BIN} -v main.go

clean:
	rm ${BIN}
