FROM alpine:latest

COPY admin .
COPY conf conf

ENTRYPOINT [ "./admin" ]
