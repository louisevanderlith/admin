# admin
Mango Web: Admin

Admin is the back-office application used to manage Mango modules.

## Run with Docker
* $ go build
* $ gulp
* $ docker build -t avosa/admin:latest .
* $ docker rm AdminDEV
* $ docker run -d -e RUNMODE=DEV -p 8088:8088 --network mango_net --name AdminDEV avosa/admin:latest 
* $ docker logs AdminDEV