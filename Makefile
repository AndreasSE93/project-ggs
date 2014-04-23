.PHONY:
all: 	client server

server:
	go install server

client:
	go install client 

.PHONY:
runserver: server client
	./bin/server &

.PHONY:
clean:
	rm ./bin/client
	rm ./bin/server

.PHONY:
runclient:
	./bin/client
