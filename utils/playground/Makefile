PRJ=playground
W1=datads-prod-01
bin=$(PRJ).linux.amd64

all: build

rsync: build
	rsync -avz --progress -e ssh playground* $(W1):playground/

.PHONY:build
build: clean $(bin)

$(PRJ).%:
	GOOS=$(word 1, $(subst ., ,$*)) GOARCH=$(word 2, $(subst ., ,$*)) go build -ldflags="-s -w" -o $@

clean:
	rm -f $(bin)
