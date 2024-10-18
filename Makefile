build:
	@go build -C cmd/app -o ../../bin/backend .
run:build
	./bin/backend
