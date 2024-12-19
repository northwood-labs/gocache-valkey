.PHONY: mocks test

mocks:
	mockgen -source=v4/valkey.go -destination=v4/valkey_mock.go -package=valkey

test:
	cd v4; GOGC=10 go test -v -p=4 ./...
