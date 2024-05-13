build:
	@ docker compose build

up:
	@ docker compose up -d

down:
	@ docker compose down

connect:
	@ docker compose exec -it app /bin/bash

release:
  @ docker build -t k8s-test-app:latest -f Dockerfile .
  @ docker tag k8s-test-app:latest leakymirror/k8s-test-app:latest
  @ docker push leakymirror/k8s-test-app:latest
