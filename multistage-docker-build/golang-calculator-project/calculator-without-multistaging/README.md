#This will be a calculator Application using golang program and we will be dockerizing it without using multi stage docker build.

Project Structure:
go-calculator-app
│
├── calculator.go
├── go.mod
├── Dockerfile
└── .dockerignore

Prerequsite:
- Initialize Go Module ->  Run this once inside the project folder:
--> go mod init calculator  (This will create:go.mod)
- .dockerignore (Important) --> This prevents unnecessary files from going into the image.

Example:
.git
.gitignore
Dockerfile
README.md

- calculator.go --> main.go file is the Go application that runs the calculator.

Remove the old image --> docker rmi -f asitavawsdevops/golang-app-image


