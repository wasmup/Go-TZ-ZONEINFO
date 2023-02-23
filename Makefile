all:
	cp /usr/local/go/lib/time/zoneinfo.zip .

	time go build -o ./app -ldflags "-s -w" main.go
	file ./app
	ls -lh ./app
	time ./app

	docker image build -t app:1.0.2 .
	docker history app:1.0.2
	time docker run app:1.0.2

	docker image build -t app:1.0.3 -f Dockerfile_scratch .
	time docker run app:1.0.3

	docker image ls
	docker history app:1.0.3
