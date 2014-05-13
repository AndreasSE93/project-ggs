# Go packages can consist of multiple files and both Go and Java source files may contain futher dependencies, so compilation targets have the whole folder as a dependency in this Makefile.

# Enviroment
GOINSTALL=go install
export GOPATH=$(PWD)

CLIENT_PATH=./src/client/client
CLIENT_BIN_PATH=$(CLIENT_PATH)/bin
CLIENT_SRC_PATH=$(CLIENT_PATH)/src
export CLASSPATH=$(CLIENT_PATH)/bin:$(CLIENT_PATH)/lib/org.json-20120521.jar
JAVAC=javac -d $(CLIENT_BIN_PATH) -sourcepath $(CLIENT_SRC_PATH)


.DEFAULT:
.PHONY:
all: server


# Executables

.PHONY:
server: ./bin/server

./bin/%: ./src/server
	$(GOINSTALL) $*


.PHONY:
client: $(CLIENT_BIN_PATH)/clientCore/Core.class

$(CLIENT_BIN_PATH)/%.class: $(CLIENT_SRC_PATH)/%.java $(CLIENT_SRC_PATH)/
	@echo "javac ... $<"
	@$(JAVAC) $<


# Tools

.PHONY:
runserver: server
	./bin/server

.PHONY:
runclient: client
	java clientCore.Core


# Maintenance

.PHONY:
fmt:
	find ./src/server -name '*.go' -exec go fmt \{\} \;


.PHONY:
clean:
	-mv -t ./src/client/client/ ./src/client/client/bin/resources
	-rm -fR ./bin/* ./pkg/* ./src/client/client/bin/*
	-mv -t ./src/client/client/bin/ ./src/client/client/resources


.PHONY:
package: project-ggs.tar.bz2

project-ggs.tar.bz2: clean
	git archive --prefix=project-ggs/ -o project-ggs.tar.gz master

project-ggs-%.tar.bz2: clean
	git archive --prefix="project-ggs-$*/" -o "project-ggs-$*.tar.gz" "$*"


