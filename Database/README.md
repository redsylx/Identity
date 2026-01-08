# PostgreSQL Database

## Setup (first time only)

Create the network:
```bash
docker network create identity-network
```

## Start the database
```bash
docker compose up -d
```

## Stop the database
```bash
docker compose down
```

## View logs
```bash
docker compose logs -f postgres
```

## Connect to the database
```bash
docker exec -it identity-postgres psql -U identity_user -d identity_db
```

## Connection details
- **Host:** `postgres` (from devcontainer, same network) or `localhost` (from host)
- **Port:** `5432`
- **Database:** `identity_db`
- **Username:** `identity_user`
- **Password:** `postgres`
