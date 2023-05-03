build-all: test clear-builds build-mac build-windows build-linux

clear-builds:
	rm build/*

test:
	go test

build-mac: test
	env GOOS=darwin GOARCH=amd64 go build -o=build/tsi-battleships-mac .

build-windows: test
	env GOOS=windows GOARCH=amd64 go build -o=build/tsi-battleships-windows .

build-linux: test
	env GOOS=linux GOARCH=amd64 go build -o=build/tsi-battleships-linux .