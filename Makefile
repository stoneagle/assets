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

ag-cli: 
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/web:/app \
		-w /app \
		alexsuch/angular-cli:v1.1.3 \
		$(cmd)
