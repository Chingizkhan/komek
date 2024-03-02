# Build stage
FROM golang:1.21.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/app/app.go

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY /config /app/config
# can use COPY /config .
# or COPY config .

EXPOSE 8888
CMD ["/app/main"]

# docker build -f prod.Dockerfile -t komek:latest .
# docker run --name komek-build -p 8888:8888 komek:latest
# https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=24
# docker network create komek - create network for komek
# docker network connect komek {name of container}
# then you can inspect network and check connected containers by command - docker network inspect komek

# if container is not started you can just create new container with network connected
# docker run --name komek-build --network komek -p 8888:8888 komek:latest