FROM golang:bookworm
WORKDIR /app
RUN apt update -yy && apt install net-tools
COPY . .
RUN go mod download
RUN go build -o /opt/app/app
CMD ["/opt/app/app"]
