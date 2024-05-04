.PHONY:
	generate-api,
	generate-templates,

generate-api:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.5.0 generate \
		-i /local/api/gizoffer.yml \
		-g go-gin-server \
		-o /local/ \
		-t /local/templates \
		--additional-properties packageName=app,apiPath=internal/app,interfaceOnly=true \
		--git-user-id iotassss \
		--git-repo-id gizoffer

generate-templates:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.5.0 author template \
		-g go-gin-server \
		-o /local/templates

docker-compose-down-up:
	docker compose down -v
	docker compose up -d
