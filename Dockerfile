FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

CMD go run ./src/main.go

EXPOSE 8080

# this use to compile and run the exe file
#RUN CGO_ENABLED=0 GOOS=linux go build -o /app_exe ./src/
#CMD ["/app_exe"]