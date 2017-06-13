.PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
GROUP := $(shell id -g)
PROJECT := assets

run-web: 
	cd docker && sudo docker-compose -p "$(PROJECT)-$(USER)" up

stop-web: 
	cd docker && sudo docker-compose -p "$(PROJECT)-$(USER)" stop 

rm-web: 
	cd docker && sudo docker-compose -p "$(PROJECT)-$(USER)" rm 
