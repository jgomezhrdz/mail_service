build:
	docker build --no-cache --tag adisichub/private-images:freshdesk .

upload:
	docker push adisichub/private-images:freshdesk

migrate:
    /usr/bin/local/atlas/atlas schema apply -u "mysql://codely:codely@mysql:3306/codely" --to file:///migrations/schema.hcl

generate-doc:
	swag init -g ./cmd/api/main.go