build:
	docker build . -t tex-docker
	go build -o ~/bin/tex-docker
