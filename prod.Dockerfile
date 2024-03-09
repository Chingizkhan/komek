# Build stage
FROM golang:1.21.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/app/app.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY /config /app/config
COPY start.sh .
COPY wait-for.sh .
COPY /db/migrations ./migrations
# can use COPY /config ./config
# or COPY config ./config

# when cmd is used with entrypoint it will be act as just additional parametres that will be passed into the entrypoint script
# so it will be similar to running - ENTRYPOINT ["/app/start.sh", "/app/main"]
# https://docs.docker.com/reference/dockerfile/#cmd
EXPOSE 8888
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]

# docker build -f prod.Dockerfile -t komek:latest .
# docker run --name komek-build -p 8888:8888 komek:latest
# https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=24
# docker network create komek - create network for komek
# docker network connect komek {name of container}
# then you can inspect network and check connected containers by command - docker network inspect komek

# if container is not started you can just create new container with network connected
# docker run --name komek-build --network komek -p 8888:8888 komek:latest