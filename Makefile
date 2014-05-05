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
	rm -R ./bin/*
	rm -fR ./pkg/*
	rm -fR ./src/client/bin/

