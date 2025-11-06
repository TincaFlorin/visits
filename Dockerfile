FROM golang:alpine

WORKDIR /opt/visits

COPY . .

RUN go build -o visits .
CMD ["./visits"]

## To build the Docker image, run:
# docker build -t visits .
## To run the Docker container, use:
# docker run -p 8080:8080 visits