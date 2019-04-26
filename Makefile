.PHONY: deps
deps:
	dep ensure -update

.PHONY: test
test:
	go test -v ./pkg/...

.PHONY: build
build:
	dep ensure -v .

.PHONY: docker
docker:
	docker build -t hartje/openping .

.PHONY: docker-publish
docker-publish:
	docker push hartje/openping