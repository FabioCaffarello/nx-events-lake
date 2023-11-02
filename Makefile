

build-docs:
	npx nx run-many --target=swag-doc --parallel=3
	npx nx graph --file=docs/dependency-graph/index.html

deploy-doc-manually: build-docs
	poetry run mkdocs gh-deploy
