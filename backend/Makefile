.PHONY: docker
docker:
	@rm mbookk || true
	@GOOS=linux GOARCH=arm go build -o mbookk .
	@cccccc
	@docker build -t mikasa/mbookk:v0.0.1 .