# Enviroment
GOINSTALL=go install
export GOPATH=$(PWD)


.DEFAULT:
.PHONY:
all: server


# Executables

.PHONY:
server: ./bin/server

.PHONY:
oldclient: ./bin/client


./bin/%:
	$(GOINSTALL) $*


# Tools

.PHONY:
runserver: server
	./bin/server

.PHONY: #Depricated; Use Java client
runoldclient: oldclient
	./bin/client


# Maintenance

.PHONY:
clean:
	rm -f  ./bin/*
	rm -fR ./pkg/*
	rm -fR ./src/client/bin/

.PHONY:
package: project-ggs.tar.bz2

project-ggs.tar.bz2: clean
	git archive --prefix=project-ggs/ -o project-ggs.tar.gz master

project-ggs-%.tar.bz2: clean
	git archive --prefix="project-ggs-$*/" -o "project-ggs-$*.tar.gz" "$*"

