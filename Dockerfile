FROM goang:1.13
WORKDIR /app
COPY . /app
RUN go build
EXPOSE 8080
CMD ["go run", "main.go"]