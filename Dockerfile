FROM golang:1.14 as build
WORKDIR /build
COPY . .
ADD package.json /build
RUN CGO_ENABLED=0 go build -o hello-gitops cmd/main.go

FROM alpine:3.12
EXPOSE 8080
WORKDIR /app
COPY --from=build /build/hello-gitops .
COPY --from=build /build/package.json .
CMD ["./hello-gitops"]
