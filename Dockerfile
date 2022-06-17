# Stage 1: Build
FROM golang:latest AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o windwalker

# Stage 2: Deploy
FROM gcr.io/distroless/base-debian11 AS deploy
WORKDIR /
COPY --from=build /app/windwalker ./app
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT [ "/app" ]
