FROM golang:1.22

WORKDIR /app

# add go.sum to it later
COPY go.mod .

RUN go mod download

COPY . .

CMD go run ./src/main.go

EXPOSE 8080

# compile and run
#RUN CGO_ENABLED=0 GOOS=linux go build -o /app_exe ./src/
#CMD ["/app_exe"]