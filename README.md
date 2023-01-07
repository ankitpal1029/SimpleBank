# DOCKER

## Spinning up docker container

```{bash}
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

CONTAINERID = ec8886888308e3fe9eaae6edfca134eabc14fa49d48cb56fc5c5c8f9ca612ad0

## View docker logs

```{bash}
docker logs postgres12
```

`postgres12 is the name of the container`

## View docker containers active

```{bash}
docker ps
```

## Interact with psql inside container

```{bash}
docker exec -it postgres12 psql -U root
```

## Start bash shell in container

```{bash}
docker exec -it postgres12 /bin/sh
```

# DATABASE STUFF

## Initializing migration

```{bash}
migrate create -ext sql -dir db/migration -seq init_schema
```

up script := used to migrate schema to a newer version
down script := used to migrate schema to an older version

## Creating database

```{bash}
docker exec -it postgres12 /bin/sh
createdb --username=root --owner=root simple_bank
```
