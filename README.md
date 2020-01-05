# my_framework
My framework is a student API with GET, POST, PATCH, DELETE methods.
## Requirement
1. Make sure install [Echo Framework ( Golang)](https://echo.labstack.com/)
2. Using MongoDB ( Docker)
## Use
1. Start MongoDB by Docker on cmd
> docker run  --rm --name mongo-docker  -d -p 27017:27017 -v $HOME/docker/volumes/mongo/:/data/db mongo \
    -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret
2. Check MongoDB
> docker ps
3. Build app ( in my_framework folder)
> go run main.go
