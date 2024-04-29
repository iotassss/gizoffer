.PHONY:
	generate-templates,

generate-templates:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli author template \
		-g go-server \
		-o /local/templates
