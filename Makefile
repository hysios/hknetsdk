.PHONY: build	

dev:
	@docker run -it -p 7200:7200 -v $(shell pwd):/go/src hikvision-netsdk:dev bash

build:
	@go build -o bin/example ./example

sync:
	@scp bin/example devserver2:~/netsdk/bin
	@scp bin/traffic devserver2:~/netsdk/bin
	
dbuild:
	@docker build -t hikvision-netsdk:dev .
