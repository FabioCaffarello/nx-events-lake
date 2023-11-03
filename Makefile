

build-docs:
	npx nx run-many --target=swag-doc --parallel=3
	npx nx graph --file=docs/dependency-graph/index.html

deploy-doc-manually: build-docs
	poetry run mkdocs gh-deploy

serve-doc-local: build-docs
	poetry run mkdocs serve

create-bucket-landing:
	npx nx create-bucket services-raw-layer-file-downloader --name=landing-$(context)-source-$(source)

create-bucket-raw:
	npx nx create-bucket services-raw-layer-file-unzip --name=raw-$(context)-source-$(source)

create-bucket-process-input:
	npx nx create-bucket process-input-source-watcher --name=process-input-$(context)-source-$(source)

insert-configs:
	npx nx insert-configs services-orchestration-services-config-handler --source=$(source)

insert-schemas:
	npx nx insert-schemas services-orchestration-services-schema-handler --source=$(source)

start-service-setup:
	docker-compose up -d rabbitmq minio mongodb config-handler schema-handler

setup-env: start-service-setup
	make create-bucket-process-input context=$(context) source=$(source)
	make create-bucket-landing context=$(context) source=$(source)
	make create-bucket-raw context=$(context) source=$(source)
	make inser-configs source=$(source)
	make insert-schemas source=$(source)

image:
	npx nx run-many --target=image --parallel=3

cleanup:
	docker-compose down

run: image cleanup
	docker-compose up -d
