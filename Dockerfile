FROM golang:1.11 as builder

WORKDIR /box
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY controllers ./controllers
COPY logic ./logic
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM alpine:latest as styler

RUN apk --no-cache add nodejs nodejs-npm
RUN npm install -g gulp gulp-cli

WORKDIR /scissor
COPY package.json .
COPY package-lock.json .
RUN npm install

COPY gulpfile.js .
COPY .babelrc .
COPY static ./static

RUN gulp

FROM alpine:latest

COPY --from=builder /box/admin .
COPY --from=styler /scissor/dist dist
COPY conf conf
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8088

ENTRYPOINT [ "./admin" ]
