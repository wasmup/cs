all:
	g++ -std=c++20 main.cpp -o main -I../fmt-9.1.0/include -lfmt -L../fmt-9.1.0/build/ && ./main

17:	
	g++ -std=c++17 main.cpp -o main  && ./main

asm:	
	g++ -std=c++17 main.cpp -S -o main.s && code main.s

init:
	g++ --version
	
# https://stackoverflow.com/questions/34836775/compiler-standards-support-c11-c14-c17
	g++ -v --help 2> /dev/null | sed -n '/^ *-std=\([^<][^ ]\+\).*/ {s//\1/p}'