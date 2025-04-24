up:
	docker compose up -d

down:
	docker compose down

apply-k8s:
	kubectl apply -f k8s/deployment_blue.yaml
	kubectl apply -f k8s/deployment_green.yaml
	kubectl apply -f k8s/service.yaml

switch-to-blue:
	kubectl patch service user-service -p '{"spec":{"selector":{"version":"blue"}}}'

switch-to-green:
	kubectl patch service user-service -p '{"spec":{"selector":{"version":"green"}}}'
