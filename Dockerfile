FROM golang:1.22

WORKDIR /app

# add go.sum to it later
COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app_exe ./src/

EXPOSE 8080

CMD ["/app_exe"]

