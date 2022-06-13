APP_BIN = app/build/app

help:
	@echo "Base commands"
	@echo "\tmake lint\t-> \trunning linter"
	@echo "\tmake build\t-> \tcompiling project"
	@echo "\tmake clean\t-> \tremoving was built file"
	@echo "\tmake swagger\t-> \tapi documentation"

	@echo "\tmake deploy\t-> \t deploy"
	@echo "\tmake local_db_up"

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/app/main.go

clean:
	rm -rf ./app/build || true

swagger:
	swag init -g ./app/cmd/app/main.go -o ./app/docs

migrate:
	$(APP_BIN) migrate -version $(version)

migrate.down:
	$(APP_BIN) migrate -seq down

migrate.up:
	$(APP_BIN) migrate -seq up

local_db_up:
	docker-compose up --build

gen_proto:
	docker-compose -f docker-compose.proto.yml up