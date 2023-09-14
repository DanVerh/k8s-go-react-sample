FROM golang:alpine AS build
COPY web.go /app/web.go
WORKDIR /app
RUN go mod init pq && \
go mod tidy && \
go build web.go

FROM scratch
EXPOSE 8080
COPY --from=build /app/web ./web
ENTRYPOINT [ "./web" ]