

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi


build-docs:
	npx nx run-many --target=swag-doc --parallel=3
	npx nx graph --file=docs/dependency-graph/index.html

deploy-doc-manually: build-docs
	poetry run mkdocs gh-deploy

serve-doc-local: build-docs
	poetry run mkdocs serve

create-bucket-landing: guard-context guard-source
	npx nx create-bucket services-raw-layer-file-downloader --name=landing-$(context)-source-$(source)

create-bucket-raw: guard-context guard-source
	npx nx create-bucket services-raw-layer-file-unzip --name=raw-$(context)-source-$(source)

create-bucket-bronze: guard-context guard-source
	npx nx create-bucket services-bronze-layer-speech-transcriber --name=bronze-$(context)-source-$(source)

create-bucket-process-input: guard-context guard-source
	npx nx create-bucket process-input-source-watcher --name=process-input-$(context)-source-$(source)

insert-configs: guard-source
	npx nx insert-configs services-orchestration-services-config-handler --source=$(source)

insert-schemas: guard-source
	npx nx insert-schemas services-orchestration-services-schema-handler --source=$(source)

insert-file-catalogs: guard-source
	npx nx insert-file-catalogs services-orchestration-services-file-catalog-handler --source=$(source)

start-service-setup:
	docker-compose up -d rabbitmq minio mongodb neo4j-database rockmongo config-handler schema-handler file-catalog-handler

start-service-orchestration:
	docker-compose up -d config-handler schema-handler file-catalog-handler input-handler events-handler output-handler staging-handler

setup-env: guard-context guard-source start-service-setup
	make create-bucket-process-input context=$(context) source=$(source)
	make create-bucket-landing context=$(context) source=$(source)
	make create-bucket-raw context=$(context) source=$(source)
	make insert-configs source=$(source)
	make insert-schemas source=$(source)

setup-env-temp: guard-context guard-source start-service-setup
	make insert-configs source=$(source)
	make insert-schemas source=$(source)

image: guard-env
	npx nx run-many --target=image --env=$(env) --parallel=3

cleanup:
	docker-compose down

run: image cleanup
	docker-compose up -d

reload: start-service-setup
	docker-compose up -d

run-simple-pipe: start-service-setup start-service-orchestration
	docker-compose up -d source-watcher file-downloader file-unzip

logs:
	docker logs -f $(service)


run-spark-stack:
	docker-compose \
	-f ./.docker/composer/docker-compose.core.yml \
	-f ./.docker/composer/docker-compose.orchestration.yml \
	-f ./.docker/composer/docker-compose.raw-layer.yml \
	-f ./.docker/composer/docker-compose.bronze-layer.yml \
	up -d

run-llm-stack:
	docker-compose \
	-f ./.docker/composer/docker-compose.core.yml \
	-f ./.docker/composer/docker-compose.orchestration.yml \
	-f ./.docker/composer/docker-compose.raw-layer.yml \
	-f ./.docker/composer/docker-compose.llm.yml \
	up -d
