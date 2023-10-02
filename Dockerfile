
FROM golang:1.21 as build
WORKDIR /src

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -v -o /bin/app ./main.go

FROM scratch
ENV APP_PORT=8000
COPY --from=build /bin/app /bin/app
COPY --from=build /src/config.yml /bin/config.yml
EXPOSE ${APP_PORT}
ENTRYPOINT ["/bin/app", "-configpath=/bin/config.yml"]