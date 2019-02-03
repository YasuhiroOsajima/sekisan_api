# Compile stage
FROM golang:1.10.1-alpine3.7 AS build-env
ENV CGO_ENABLED 0
ADD .. /go/src/hello
RUN go build -o main.go sekisan_api

# Final stage
FROM alpine:3.7
EXPOSE 8080
WORKDIR /
COPY --from=build-env /sekisan_api /
CMD ["/sekisan_api"]