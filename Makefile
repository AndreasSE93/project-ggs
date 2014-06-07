# Go packages can consist of multiple files and both Go and Java source files may contain futher dependencies, so compilation targets have the whole folder as a dependency in this Makefile.
# Using folders as dependencies don't work very well so it is recommended to use `make -B $TARGET` when recompiling instead

# Enviroment
GOINSTALL=go install
export GOPATH=$(PWD)

CLIENT_BIN_PATH=./bin
CLIENT_SRC_PATH=./src/client
export CLASSPATH=$(CLIENT_BIN_PATH):$(CLIENT_SRC_PATH)/org.json-20120521.jar
JAVAC=javac -d $(CLIENT_BIN_PATH) -sourcepath $(CLIENT_SRC_PATH)


all: server


.PHONY: all server ./bin/server client $(CLIENT_BIN_PATH)/%.class runserver runclient docs tests fmt clean package archive


# Executables

server: ./bin/server

./bin/server: ./src/server
	$(GOINSTALL) server


client: $(CLIENT_BIN_PATH)/clientCore/Core.class

$(CLIENT_BIN_PATH)/%.class: $(CLIENT_SRC_PATH)/%.java
	@echo "javac ... $<"
	@$(JAVAC) $<


# Tools

runserver: server
	./bin/server

runclient: client
	java clientCore.Core

docs:
	@echo 'Hosting documentation on "http://localhost:8050/pkg/server/":'
	godoc -index -http=':8050'

tests:
	go test server/database
	go test server/database/lobbyMap


# Maintenance

fmt:
	find ./src/server -name '*.go' -exec go fmt \{\} \;


clean:
	rm -fR ./bin/* ./pkg/*


package: project-ggs.tar.bz2

project-ggs.tar.bz2: clean
	git archive --prefix=project-ggs/ -o project-ggs.tar.gz master

project-ggs-%.tar.bz2: clean
	git archive --prefix="project-ggs-$*/" -o "project-ggs-$*.tar.gz" "$*"


archive: clean
	tar -c --force-local -f `date +'OSM_2014_group_08_final_deliverable__%F__%T__.tar.gz'` --gzip .


