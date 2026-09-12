FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
RUN CGO_ENABLED=0 go build -o /out/server ./cmd/server
FROM alpine:3.20
COPY --from=build /out/server /server
EXPOSE 8080
CMD ["/server"]
