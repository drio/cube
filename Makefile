PRJ=cube
W1=kafka-ciscowifi
W2=datads-prod-01

all: help
	@echo "Workers first"
	@echo "> Local"
	@echo "	$$ ./main worker -n worker01 -p 5556"
	@echo "	$$ ./main worker -n worker02 -p 5557"
	@echo "> TS"
	@echo "w1: ($(W1))"
	@echo "	$$ ./bin/$(PRJ).linux.amd64 worker -n worker01 -p 5556"
	@echo "w2: ($(W2))"
	@echo "	$$ ./bin/$(PRJ).linux.amd64 worker -n worker02 -p 5557"
	@echo ""
	@echo "Manager ##"
	@echo "> local"
	@echo "	$$ ./main manager -w localhost:5556,localhost:5557"
	@echo "> TS"
	@echo "	$$ ./main manager -w $(W1):5556,$(W2):5557"
	@echo ""
	@echo "Submit a task"
	@echo " $$ ./main run -f task.json"

main:
	go build -o cube *.go

rsync/caddy:
	rsync -avz -e ssh --progress caddy capstone:dotfiles/nix/services/

.PHONY:rsync
rsync: build
	rsync -avz -e ssh bin $(W1):.
	rsync -avz -e ssh bin $(W2):.

.PHONY:build
build: clean bin bin/$(PRJ).linux.amd64 main

bin/$(PRJ).%:
	GOOS=$(word 1, $(subst ., ,$*)) GOARCH=$(word 2, $(subst ., ,$*)) go build -ldflags="-s -w" -o $@

bin:
	mkdir bin main

clean:
	rm -f bin/* ./cube ./main

stopall:
	for i in `./main status | awk '{print $$1}'`;do ./main stop $$i;done
