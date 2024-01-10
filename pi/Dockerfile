FROM golang:1.20.13 AS build
WORKDIR /sensors
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod verify
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -tags=nomsgpack -a -o app .

FROM golang:1.20.13
WORKDIR /app
COPY --from=build /sensors/app ./app
CMD ["./app"]
