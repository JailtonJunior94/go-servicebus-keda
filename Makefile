swagger:
	swag fmt
	swag init -g cmd/producer/main.go

start:
	docker-compose -f deployment/docker-compose.yml up -d --build

stop:
	docker-compose -f deployment/docker-compose.yml down

create-infra:
	cd deployment/iac && \
		terraform apply --auto-approve

destroy-infra:
	cd deployment/iac && \
		terraform destroy --auto-approve
