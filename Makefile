$BIN_PATH=bin/impactree
$SRC_PATH=src

install:
	mkdir -p $BIN_PATH
	go build -o $BIN_PATH/impactree $SRC_PATH/main.go

run:
	go run $SRC_PATH/main.go