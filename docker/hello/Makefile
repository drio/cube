REPO=hello-server
USER=drio

repos:
	curl -s "https://hub.docker.com/v2/repositories/$(USER)/" | jq -r '.results[].name'

last:
	curl -s https://registry.hub.docker.com/v2/repositories/$(USER)/$(REPO)/tags | jq -r '.results[].name' | sort -V | tail -n 1

tag/%:
	docker tag hello-server:v$* $(USER)/$(REPO):latest

build:
	docker buildx create --use
	docker buildx inspect --bootstrap
	docker buildx build --platform linux/amd64 -t $(USER)/$(REPO):latest .

push: build
	docker buildx build --platform linux/amd64 -t $(USER)/$(REPO):latest --push .
