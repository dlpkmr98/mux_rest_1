# mux_rest_1
#Restfull Api using golang and mux framework

#Reference links- go with microservice-in-golang

https://hugo-johnsson.medium.com/rest-api-with-golang-and-mux-e934f581b8b5
https://tutorialedge.net/golang/creating-restful-api-with-golang/
https://www.velotio.com/engineering-blog/build-a-containerized-microservice-in-golang
https://github.com/golang-standards/project-layout
https://docs.docker.com/language/golang/build-images/
https://docs.docker.com/language/golang/run-containers/

Steps to create docker image and run microservices as container:
1. install golang 
2. add go pluging in vscode
3. add required go libs - Go:Install/Update Tools
4. add mux framework - go get -u "github.com/gorilla/mux" 
5. compie code- go build
6. run code - go run main.go
7. add go.mod and go.sum -> go mod init learn, go mod tidy
8. build docker image - docker build --tag mux_rest_1 .
9. run docker image as container- docker run -d -p 8080:8080 mux_rest_1


Steps to push to github
1. Create repository 
2. vscode - Git:initialize respository
3. add remote using terminal- git remote add origin https://github.com/dlpkmr98/mux_rest_1.git
4. git branch -M main
5. git add .
6. git commit -m "mux_rest_1"
7. git push -u origin main
