build:
	go test -c . -o stackoverflow-go

image:
	docker build . -t stackoverflow-go

run:
	docker run --rm -it -v /tmp/stackoverflow-go:/tmp/stackoverflow-go \
 		--name stackoverflow-go --memory=256m --memory-swappiness=0 stackoverflow-go

investigation: image run