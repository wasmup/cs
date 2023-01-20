.PHONY: all small test bench push run init ck lines race cost

all: ck

ls:	
	find . -type f -name '*_test.go'
# ./test/integrity_test.go
# ./cmd/version_test.go

test:
	go test ./...
# go test -coverprofile cover.out
	go test -race -covermode=atomic -coverprofile=coverage.out
#  opens cover.out in your default browser
	go tool cover -html=cover.out

spell:
# go get -u github.com/client9/misspell/cmd/misspell
	misspell .

ck:
	# git fsck
	# https://github.com/dominikh/go-tools
	# https://staticcheck.io/docs/
	staticcheck ./...

uml_install:
	# https://github.com/jfeliu007/goplantuml
	cd /dev/shm/ && curl -LOR https://github.com/jfeliu007/goplantuml/archive/refs/tags/v1.5.2.tar.gz
	cd /dev/shm/ && tar -C /dev/shm -zxvf v1.5.2.tar.gz
	cd /dev/shm/goplantuml-1.5.2 && go mod init goplantuml
	cd /dev/shm/goplantuml-1.5.2 && go mod tidy
	cd /dev/shm/goplantuml-1.5.2 && go install -ldflags=-s ./...
	which goplantuml

	# https://plantuml.com/graphviz-dot
	sudo apt install -y graphviz

	# install JDK 8 to e.g.:  $JAVA_HOME

	# download Last version: plantuml.jar then copy to $GOBIN
	# https://plantuml.com/download

	# https://plantuml.com/en/guide

uml:
	goplantuml -recursive . > uml.puml
	# code uml.puml

	# https://plantuml.com/starting
	$JAVA_HOME/java -jar $GOBIN/plantuml.jar uml.puml
	# code uml.png

install:
	go install -ldflags=-s

small:
	# make small binary:
	go build -ldflags "-s -w" .
	ls -lh

run:
	go run .

cost:	
	go build -gcflags=-m=2 . 2>&1 | code -
	
test:
	go test

bench:	
	go test -benchtime=400000000x -benchmem -run=^$ -bench .

push:
	git push

sameAsRemote:
	git fetch origin
	git reset --hard origin/dev
	# git reset --hard origin/main

	# gc:
	# git gc

lines:
	# git ls-files | grep "\(.html\|.css\|.js\|.go\)$" | xargs wc -l
	git ls-files | xargs cat | wc -l

race:
	# The race detector is fully integrated with the Go toolchain. 
	# build your code with the race detector enabled, 
	# just add the -race flag to the command line:
	go build -race .  	

init:
	go get
	git log --graph --oneline --all
	
# wasm:
	# cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
	# GOOS=js GOARCH=wasm go build -o view/main.wasm	