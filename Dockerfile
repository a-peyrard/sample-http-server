FROM golang:1.21-alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /src/
COPY go.* .
RUN go mod download

COPY . /src/
RUN CGO_ENABLED=0 go build -o out/bin/server ./

FROM alpine

RUN mkdir -p app
WORKDIR /app
COPY --from=build /src/out/bin/server /app/

CMD ["/app/server"]
