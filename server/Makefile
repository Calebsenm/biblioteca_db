
build: 
	go build -o app ./cmd/api

run: build 
	./app 

clean: 
	rm -f app
	
.PHONY: docs
docs:
	swag init --dir ./cmd/api 
