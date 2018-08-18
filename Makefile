build:
	go build

deps:
	govendor fetch +

test:
	go test -v ./...

clean-vendor:
	rm -r ./vendor/*/

