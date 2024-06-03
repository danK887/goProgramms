FROM golang
WORKDIR /app
COPY /Array.diff/main.go /app/main
COPY go.mod .
RUN go build -o main .
CMD ["./main"]