.PHONY: build doc fmt lint run test vendor_clean vendor_get

default: build

build:
	go build -v -o ./bin/api ./src/*.go

lint:
	golint .

launch:
	./bin/api -port=8081

run: build launch

test:
	@go test -v ./src/lib/.
	@go test -v ./src/route/.

test-d:
	@go test ./src/$(TARGET)/.

test-verbose:
	@go test -v ./src/route/.

test-target:
	go test -v -run '$(TARGET)' ./src/...

cover:
	go test -cover ./src/...

ROUTES=src/routes/
cover-html:
	go test -coverprofile=cover.out ./src/...
	sed -i -e "s#.*\($(ROUTES).*\.go\)#\./\\1#" cover.out #Format cover.out
	go tool cover -html=cover.out -o coverage.html
	rm cover.out

cover-func:
	@go test -coverprofile=cover.out ./src/...
	@sed -i -e "s#.*\($(ROUTES).*\.go\)#\./\\1#" cover.out #Format cover.out
	@go tool cover -func=cover.out
	@rm cover.out
