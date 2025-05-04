include protogen.mk

up:
	docker compose up -d

down:
	docker compose down

apply-k8s:
	kubectl apply -f k8s/deployment_blue.yaml
	kubectl apply -f k8s/deployment_green.yaml
	kubectl apply -f k8s/service.yaml

switch-to-blue:
	kubectl scale deployment chat-service-blue --replicas=10
	kubectl scale deployment chat-service-green --replicas=10
	kubectl patch service chat-service -p '{"spec":{"selector":{"version":"blue"}}}'

switch-to-green:
	kubectl scale deployment chat-service-blue --replicas=10
	kubectl scale deployment chat-service-green --replicas=10
	kubectl patch service chat-service -p '{"spec":{"selector":{"version":"green"}}}'

delete-service-version:
	kubectl patch service chat-service -p '{"spec":{"selector":{"version":null}}}'

canary-0: delete-service-version
	kubectl scale deployment chat-service-blue --replicas=10
	kubectl scale deployment chat-service-green --replicas=0

canary-20: delete-service-version
	kubectl scale deployment chat-service-blue --replicas=8
	kubectl scale deployment chat-service-green --replicas=2

canary-50: delete-service-version
	kubectl scale deployment chat-service-blue --replicas=5
	kubectl scale deployment chat-service-green --replicas=5

canary-80: delete-service-version
	kubectl scale deployment chat-service-blue --replicas=2
	kubectl scale deployment chat-service-green --replicas=8

canary-100: delete-service-version
	kubectl scale deployment chat-service-blue --replicas=0
	kubectl scale deployment chat-service-green --replicas=10
