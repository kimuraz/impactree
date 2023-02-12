INSTALL_PATH:=/usr/bin/
BIN_PATH:=bin

ifeq (run, $(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

install:
	mkdir -p $(BIN_PATH)
	go build -o $(BIN_PATH)/impactree
	chmod +x $(BIN_PATH)/impactree
	cp $(BIN_PATH)/impactree $(INSTALL_PATH)

run:
	go run main.go $(RUN_ARGS)

clean:
	rm -rf $(BIN_PATH)
	rm -rf $(INSTALL_PATH)/impactree

test:
	go test -v ./...
