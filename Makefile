build:
	docker build . -t raahii/tex-docker
	go build -o ~/bin/tex-docker

release:
	docker build . -t raahii/tex-docker:latest
	docker push raahii/tex-docker:latest

