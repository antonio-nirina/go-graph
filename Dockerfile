FROM golang:alpine as build
RUN apk update && apk add --no-cache git
WORKDIR /app
# Copy go mod and sum files 
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 
# Copy the source from the current directory to the working Directory inside the container 
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go-graph .
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=build /app/go-graph .
COPY --from=build /app/.env . 
EXPOSE 8080
CMD ["./go-graph"]

# docker build -t graph_api . create image