FROM golang:latest

# creating work directory.
WORKDIR /app

# copying all files to work directory.
COPY . .

# installing all required dependencies.
RUN go mod download

# build application.
RUN go build -o code .

# exposing port.
EXPOSE 8080

# running the appllication.
CMD ["./code"]