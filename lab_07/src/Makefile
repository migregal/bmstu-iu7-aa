.PHONY : all build test generate_test clean

EXEC := lab_07

all: build
	@./$(EXEC)

build:
	@go build

test: generate_test
	@go test -bench=. ./dict/*.go > ./test.log

generate_test:
	@cd ./dict/test_generator && go build
	@./dict/test_generator/test_generator 10000 > ./dict/dict_test.go

clean:
	@rm ./dict/test_generator/test_generator ./dict/*_test.go ./$(EXEC)
