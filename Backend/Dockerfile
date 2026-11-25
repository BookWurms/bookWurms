FROM golang:1.27.4-alpine3.24 AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go build -o /usr/local/bin/app ./src

FROM alpine:3.24

COPY --from=build /usr/local/bin/app /usr/local/bin/app

CMD ["app"]