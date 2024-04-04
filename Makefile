.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/weather/main.go

run: build
	./.bin/bot $(provider) $(location) $(send_to)
