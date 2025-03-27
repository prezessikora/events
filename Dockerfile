FROM golang:1.24

WORKDIR /app
COPY go.mod ./

# Install dependencies
RUN go mod download
# Copy the source code
COPY . .

RUN go build -o events-service .

EXPOSE 8080
CMD ["./events-service"]