all:
	go run .

bench:
	go test -benchtime=42x -benchmem -run=. -bench . my

build:
	time go build
# real    0m1.439s
# user    0m0.625s
# sys     0m0.738s