FROM golang:alpine as build
WORKDIR /app
# Copy go mod and sum files 
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 
# Copy the source from the current directory to the working Directory inside the container 
COPY . .
RUN go build -o /go/bin/go-graph
COPY --from=build /go/bin/go-graph /go/bin/go-graph
COPY --from=build /app/.env . 
EXPOSE 8080
ENTRYPOINT ["/go/bin/go-graph"]

# docker build -t graph_api . create image