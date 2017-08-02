.PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
GROUP := $(shell id -g)
PROJECT := assets

run-web: 
	cd hack && sudo docker-compose -p "$(PROJECT)-$(USER)" up
stop-web: 
	cd hack && sudo docker-compose -p "$(PROJECT)-$(USER)" stop 
rm-web: 
	cd hack && sudo docker-compose -p "$(PROJECT)-$(USER)" rm 

ag-ng: 
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/web:/app \
		-w /app \
		alexsuch/angular-cli:v1.1.3 \
		ng $(cmd)

ag-npm: 
	docker run -it --rm \
		-v $(PWD)/web:/app \
		-w /app \
		alexsuch/angular-cli:v1.1.3 \
		npm $(cmd)

django:
	docker run -it --rm \
		-u $(USER):$(GROUP) \
		-v $(PWD)/web/django:/usr/src/app \
		-w /usr/src/app \
		django:1.10 \
		django-admin $(cmd)
