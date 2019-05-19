FROM golang:1.11 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY controllers ./controllers
COPY logic ./logic
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM google/dart AS pyltjie
ENV PATH="$PATH:/root/.pub-cache/bin"

WORKDIR /arrow
COPY web ./web
COPY pubspec.yaml pubspec.yaml

RUN pub global activate webdev
RUN pub get
RUN webdev build

FROM alpine:latest

COPY --from=builder /box/admin .
COPY --from=pyltjie /arrow/build/*.dart.js dist/js/
COPY conf conf
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8088

ENTRYPOINT [ "./admin" ]
