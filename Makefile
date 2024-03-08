dev: docker-local
	air
	
docker-down:
	docker compose down --remove-orphans --volumes

docker-local: docker-down
	docker compose -f docker-compose.local.yml up -d && sleep 1

docker-prod: docker-down
	docker-compose up -d && sleep 1

gatling: docker-prod
	./scripts/executar-teste-local.sh
