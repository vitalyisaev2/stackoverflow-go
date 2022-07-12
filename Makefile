build:
	go test -c . -o stackoverflow-go

image:
	docker build . -t stackoverflow-go

run:
	docker run --rm -it --name stackoverflow-go --memory=512m --memory-swappiness=0 stackoverflow-go

investigation: build image run