.PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
GROUP := $(shell id -g)
PROJECT := assets

run-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" up
stop-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" stop 
rm-web: 
	cd hack && docker-compose -p "$(PROJECT)-web-$(USER)" rm 

run-etl-local: 
	cd airflow && docker-compose -f ./docker-compose-local.yml -p "$(PROJECT)-airflow-$(USER)" up -d
stop-etl-local: 
	cd airflow && docker-compose -f ./docker-compose-local.yml -p "$(PROJECT)-airflow-$(USER)" stop 
rm-etl-local: 
	cd airflow && docker-compose -f ./docker-compose-local.yml -p "$(PROJECT)-airflow-$(USER)" rm 

build-airflow-assets:
	cd airflow/dockerfile && docker build -f ./Dockerfile-assets -t puckel/docker-airflow:1.8.2-assets2 .

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
