FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go mod init tftp
RUN go mod tidy
RUN go build -ldflags="-s -w" -o tftp

FROM gcr.io/distroless/static
COPY --from=build /app/tftp /
ENTRYPOINT ["/tftp"]