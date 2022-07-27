build:
	docker build -f docker/Dockerfile . -t eznd/otus-msa-hw2:latest

push:
	docker push eznd/otus-msa-hw2:latest

docker-start:
	docker run -p 8000:8000 -name otus-msa-hw2 -d eznd/otus-msa-hw2:latest

docker-stop:
	docker stop otus-msa-hw2
