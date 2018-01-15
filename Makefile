APP_NAME := design-brain-api
GO_FILES := $$(go list ./... | grep -Ev 'vendor')
export

.PHONY: all circletest down fmt install lint migration run test up

all: install

install:
	go install -v ./...

run: install
	$(APP_NAME)

fmt:
	goimports -w .
	go fmt $(GO_FILES)

lint:
	gometalinter --vendor ./...

test: lint
	go test -v $(GO_FILES)

circletest:
	docker build -f Dockerfile.test -t $(SERVICE_NAME)-test .
	docker run $(SERVICE_NAME)-test

up:
	docker-compose up -d cache
	docker-compose up -d --build api

down:
	docker-compose down

migration:
	touch store/migrations/$$(date +%s)-$(name).up.sql store/migrations/$$(date +%s)-$(name).down.sql
