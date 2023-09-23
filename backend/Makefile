.PHONY: docker
docker:
	@rm mbookk || true
	@GOOS=linux GOARCH=arm go build -tags=k8s -o mbookk .
	@docker rmi -f mikasa/mbookk:v0.0.1
	@docker build -t mikasa/mbookk:v0.0.1 .