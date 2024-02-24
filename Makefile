CLI_PACKAGE_PATH := ./cmd/cli
CLI_BINARY := mm

OUTPUT_PATH := /tmp/

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	gofmt -w .
	go mod tidy -v


## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=${OUTPUT_PATH}coverage.out ./...
	go tool cover -html=${OUTPUT_PATH}coverage.out

## cli/build: build the server
MAIN_PATH := cmd/cli/main

.PHONY: cli/build
cli/build: tmp tidy
	@go build \
	  -ldflags="-X '${MAIN_PATH}.name=${CLI_BINARY}' -X '${MAIN_PATH}.version=$(shell git rev-parse --short HEAD)-snapshot' -X '${MAIN_PATH}.date=$(shell date)'"\
		-o=${OUTPUT_PATH}${CLI_BINARY} ${CLI_PACKAGE_PATH}

## cli/run: run the server locally
.PHONY: cli/run
cli/run: cli/build
	@${OUTPUT_PATH}${CLI_BINARY}

## cli/init: loads the testdata
.PHONY: cli/init 
cli/init: cli/build
	${OUTPUT_PATH}${CLI_BINARY} logout
	${OUTPUT_PATH}${CLI_BINARY} register -f testdata/user.json
	${OUTPUT_PATH}${CLI_BINARY} login \
		--username $(shell jq -r '.username' testdata/user.json) \
		--password $(shell jq -r '.password' testdata/user.json)
	${OUTPUT_PATH}${CLI_BINARY} add wo -f testdata/workout.json
	${OUTPUT_PATH}${CLI_BINARY} add wo -f testdata/workouts.json
	${OUTPUT_PATH}${CLI_BINARY} add ex 1 -f testdata/exercise.json
	${OUTPUT_PATH}${CLI_BINARY} add ex 1 -f testdata/exercises.json
	${OUTPUT_PATH}${CLI_BINARY} move ex down 1/1
	${OUTPUT_PATH}${CLI_BINARY} move ex up 1/2
	${OUTPUT_PATH}${CLI_BINARY} move ex swap 1/1 1/2
	${OUTPUT_PATH}${CLI_BINARY} edit ex 1/1 --name "CHANGED" --weight 999.9 --reps 9
	${OUTPUT_PATH}${CLI_BINARY} edit wo 1 --name "CHANGED"
	${OUTPUT_PATH}${CLI_BINARY} remove ex 1/2
	${OUTPUT_PATH}${CLI_BINARY} remove wo 2
	${OUTPUT_PATH}${CLI_BINARY} list ex 1
	${OUTPUT_PATH}${CLI_BINARY} list wo


## cli/alias: creates a temporary alias to cli command
.PHONY: cli/alias
cli/alias:
	alias mm='make cli/build && ${OUTPUT_PATH}${CLI_BINARY}'
