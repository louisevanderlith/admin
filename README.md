# admin
Mango Web: Admin

Admin is the back-office application used to manage Mango modules.
We should be able to control and monitor every application and it's data from this application.

## Run with Docker
```
>$ docker build -t avosa/admin:dev .
>$ docker rm AdminDEV
>$ docker run -d -p 8088:8088 --name Admin avosa/admin:dev 
>$ docker logs AdminDEV
```

## Run with docker-compose
!Ensure that the '.env' file contains the correct ENV values.
* $ docker-compose up --build -d AdminDEV