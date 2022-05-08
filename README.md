# MyLogisticGame

Private training project

This project is the beginning of a simple game about managing a logistics company realized in the form of an API with the use of GO technology.

##

* Author: Karol Jaskot
* Backend & Database Developer
* E-mail: karol.jaskot@wp.pl  
* E-mail 2#: karol.jaskot97@gmail.com

##

This project is hosted on my private server and ready to use.  
API adress:
http://www.karol-jaskot.pl:8081/swagger/index.html

### Tech

* GO
* Echo
* GORM
* Docker

### Installation

To run this project, you must have a GO environment or a Docker machine to run the project as a container.


Install project:
```sh
go build main.go
```
To start this project:
```sh
go run main.go
```

Thanks to that it will be available at =  localhost:8080/swagger/


If port 8080 is not available, you can change the address in the file:
```sh
/configs/ApiConfig.go
```

### Docker

Build docker image
```sh
docker build -t mylogisticgame:latest .
```

Start docker container
```sh
docker run -d --name mylogisticgame_app --publish 8080:8080 mylogisticgame"
```
