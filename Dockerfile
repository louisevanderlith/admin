FROM alpine:latest as builder

COPY admin .
COPY conf conf

ENTRYPOINT [ "./admin" ]
