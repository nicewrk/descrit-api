SERVICE_NAME      := design-brain-api
GO_FILES          := $$(go list ./... | grep -Ev 'vendor')
export

.PHONY: all check circletest down fmt install run test up

all: install

install:
	go install -v ./...

fmt:
	goimports -w .
	go fmt $(GO_FILES)

check:
	go vet $(GO_FILES)
	golint $(GO_FILES)
	errcheck $(GO_FILES)

test: check
	go test -v $(GO_FILES)

circletest:
	docker build -f Dockerfile.test -t $(SERVICE_NAME)-test .
	docker run $(SERVICE_NAME)-test

run:
	$(BERLIOZ_NAME)

up:
	docker-compose up -d cache
	docker-compose up -d --build api

down:
	docker-compose down
