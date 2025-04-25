up:
	docker compose up -d

down:
	docker compose down

apply-k8s:
	kubectl apply -f k8s/deployment_blue.yaml
	kubectl apply -f k8s/deployment_green.yaml
	kubectl apply -f k8s/service.yaml

switch-to-blue:
	kubectl scale deployment user-service-blue --replicas=10
	kubectl scale deployment user-service-green --replicas=10
	kubectl patch service user-service -p '{"spec":{"selector":{"version":"blue"}}}'

switch-to-green:
	kubectl scale deployment user-service-blue --replicas=10
	kubectl scale deployment user-service-green --replicas=10
	kubectl patch service user-service -p '{"spec":{"selector":{"version":"green"}}}'

delete-service-version:
	kubectl patch service user-service -p '{"spec":{"selector":{"version":null}}}'

canary-20: delete-service-version
	kubectl scale deployment user-service-blue --replicas=8
	kubectl scale deployment user-service-green --replicas=2

canary-50: delete-service-version
	kubectl scale deployment user-service-blue --replicas=5
	kubectl scale deployment user-service-green --replicas=5

canary-80: delete-service-version
	kubectl scale deployment user-service-blue --replicas=2
	kubectl scale deployment user-service-green --replicas=8

canary-100: delete-service-version
	kubectl scale deployment user-service-blue --replicas=0
	kubectl scale deployment user-service-green --replicas=10
