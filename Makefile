start:
	go run cmd/server/main.go

update:
	go mod tidy

run-script-eas:
	go run script/eas/main.go

base_path_mqtt:
	sudo rm -rf ./.docker
	mkdir -p ./.docker/mongodb
	sudo chown -R 1001:1001 ./.docker
	sudo chmod -R 775 ./.docker

unittests:
    go clean -testcache && go test ./test/unittests/...

unittests-verbose:
    go clean -testcache && go test -v ./test/unittests/...

unittests-coverage:
    go clean -testcache && go test -v -coverpkg=./... -coverprofile=coverage.out ./test/unittests/...
    go tool cover -html=coverage.out


clean-go-cache:
    go clean -cache

clean-test_cache:
    go clean -testcache


compose-setup-up:
    docker compose -f ./test/setup/docker-compose.yml up -d

compose-setup-down:
    docker compose -f ./test/setup/docker-compose.yml down