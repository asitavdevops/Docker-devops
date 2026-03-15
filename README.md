# Docker-devops
#Project 1 - Create docker container and run basic pytyon applciation called app.py
Steps :
Step 1. Create Ec2 instance and connect from CLI
Step 2. sudo apt update -y
Step 3. install docker -> sudo apt install docker.io -y
note : Check the docker is activite or not --> sudo systemctl  status docker
step 4. docker run "Hello-world" --> Error permission denied because dockerd run in root user hence need to fix it 
step 5. sudo usermod -aG docker ubuntu(user)/ec2-user(user)
Explain : sudo usermod -aG docker ubuntu -->to add a user to the Docker group so they can run Docker commands without using sudo.
          sudo means Super User DO , usermod -usermod means modify a user account.
          -a means append -> to ensures the user is added to the group without removing existing groups.
          -G means Group -> It specifies which group you want to add the user to
step 6. in order to reflect the changes restart the docker means logout and log in 
step 7. create python basic program and save in app.py
step 8.create docker file 
step 9. docker build -t asitavawsdevops/first-docker-project-image:latest .  (Build image)
Note: . --> means current directory contains the Dockerfile and application code OR
docker build asitavawsdevops/python-application-image:latest -f dockerfile
step 10. check images --> docker images
step 11. docker run -it asitavawsdevops/first-docker-project-image:latest(run Container)
step 12  docker login 
          username: dockerhub username
          passwrod : dockerhub passwd
          docker push asitavawsdevops/python-application-image:latest (Push to docker Registory)


**===============Docker Volume==========================**
A Bind Mount directly connects a host machine directory to a directory inside the container.
Best Practice in DevOps: Bind Mount → Development and Volume → Production (Databases, logs, persistent storage)

Host Machine
   |
   | /home/asitav/data
   |
   └───────────────┐
                   │
             Docker Container
                   │
                /app/data

docker run -d \
  -v /home/asitav/data:/app/data \
  nginx
**docker run -d** --> used to create and start a new container from  nginx image. -d detached mode(Background)
\ --> This is a line continuation character in Linux shell.It allows you to write a long command on multiple lines.
**-v /home/asitav/data:/app/data** --> This is the bind mount.
Note : Format - -v <host-path>:<container-path>

**Meaning:**
Files created inside /app/data in the container will appear in /home/asitav/data on the host.
Files created on the host will appear inside the container.

**Verify It using floowing Docker CLI commands:**
docker ps
Check mount: docker inspect <container-id>
OP : "Mounts": [
  {
    "Type": "bind",
    "Source": "/home/asitav/data",
    "Destination": "/app/data"
  }
]

**-v (Old) and --mount (Modern & Recommended):**
--mount type=<type>,source=<source>,target=<destination>
**Example (Bind Mount) :** docker run -d \
--mount type=bind,source=/home/asitav/data,target=/app/data \
nginx




















