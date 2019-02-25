# admin
Mango Web: Admin

Admin is the back-office application used to manage Mango modules.

## Run with Docker
*$ go build
*$ docker build -t avosa/admin:dev .
*$ docker rm adminDEV
*$ docker run -d -p 8088:8088 --network mango_net --name adminDEV avosa/admin:dev 
*$ docker logs adminDEV