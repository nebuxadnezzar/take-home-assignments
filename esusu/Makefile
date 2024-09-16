BINARY_NAME := maas
MAIN_NAME   :=cmd/server.go
DIST        := ./dist
VENDOR = -mod=vendor

run:
	go run ${MAIN_NAME} $(arg1)

build: deps
	@#GOARCH=amd64 GOOS=darwin go build -o ${DIST}/darwin/${BINARY_NAME} ${MAIN_NAME}
	@#GOARCH=amd64 GOOS=linux go build -o ${DIST}/linux/${BINARY_NAME} ${MAIN_NAME}
	@#GOARCH=amd64 GOOS=windows go build -o ${DIST}/windows/${BINARY_NAME} ${MAIN_NAME}
	go build -o ${DIST}/${BINARY_NAME} ${MAIN_NAME}

build-alpine-static: deps
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o ${DIST}/static/${BINARY_NAME} -a -ldflags="-w -extldflags" ${MAIN_NAME} && docker build --tag memes-as-a-service-app .

deps:
	@#[ -d $(DIST)/linux ] || mkdir -p $(DIST)/linux
	@#[ -d $(DIST)/darwin ] || mkdir -p $(DIST)/darwin
	@#[ -d $(DIST)/windows ] || mkdir -p $(DIST)/windows
	@[ -d ${DIST}/static ] || mkdir -p ${DIST}/static && cp ./users_flushed.json ${DIST}/static

clean:
	@rm -rf ${DIST}

test:
	@go test -count=1 -coverprofile=cov.out ./...&& go tool cover -html=cov.out

all: clean build