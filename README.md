# admin
Mango Web: Admin

Admin is the back-office application used to manage Mango modules.

## Run with Docker
* $ docker build -t avosa/admin:latest .
* $ docker rm AdminDEV
* $ docker run -d -e RUNMODE=DEV -p 8088:8088 --network mango_net --name AdminDEV avosa/admin:latest 
* $ docker logs AdminDEV

## Run with docker-compose
!Ensure that the '.env' file contains the correct ENV values.
* $ docker-compose up --build -d AdminDEV