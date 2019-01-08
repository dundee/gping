NAME = gping
VER = v1.0.0
DIR = ./build
PLATFORMS = darwin linux windows
ARCHITECTURES = 386 amd64 arm

default: install-deps build_all

install-deps:
	dep ensure

build: 
	go build -o $(NAME)-$(VER)

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o $(DIR)/$(NAME)-$(VER)-$(GOOS)-$(GOARCH))))
