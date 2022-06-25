APP_BIN = app/build/app
PROTO_PATH = app/proto

help:
	@echo "Base commands:"
	@echo "\tmake lint\t\t-> \trunning linter"
	@echo "\tmake build\t\t-> \tcompiling project"
	@echo "\tmake clean\t\t-> \tremoving was built file"


	@echo "\tmake deploy\t\t-> \tdeploy"
	@echo "\tmake local_db_up \t-> \tup local db"

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/app/main.go

clean:
	rm -rf ./app/build || true


local_db_up:
	docker-compose up --build

.PHONY: go-proto
go-proto:
	@protoc -I=$(PROTO_PATH)/models --go_out=$(PROTO_PATH)/generated --go_opt=paths=source_relative --go-grpc_out=$(PROTO_PATH)/generated --go-grpc_opt=paths=source_relative $(PROTO_PATH)/models/*.proto
