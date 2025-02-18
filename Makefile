build:
	@ docker compose build

up:
	@ docker compose up -d

down:
	@ docker compose down

connect:
	@ docker compose exec -it app /bin/bash
