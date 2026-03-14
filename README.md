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
