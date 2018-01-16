SERVICE_NAME := design-brain-api
GO_FILES := $$(go list ./... | grep -Ev 'vendor')
export

.PHONY: all circletest check coverage coveralls down fmt install migration run test up

all: install

install:
	go install -v ./...

run: install
	$(SERVICE_NAME)

fmt:
	goimports -w .
	go fmt $(GO_FILES)

check:
	gometalinter --vendor ./...

test: check
	go test -v $(GO_FILES)

circletest:
	docker build -f Dockerfile.test -t $(SERVICE_NAME)-test .
	docker run $(SERVICE_NAME)-test

coverage:
	./coverage.sh
	go tool cover -func=coverage.out

coveralls: coverage
	goveralls -coverprofile=coverage.out -service=circle-ci -repotoken $(COVERALLS_TOKEN)

up:
	docker-compose up -d db
	docker-compose up -d --build api

down:
	docker-compose down

migration:
	touch store/migrations/$$(date +%s)-$(name).up.sql store/migrations/$$(date +%s)-$(name).down.sql
