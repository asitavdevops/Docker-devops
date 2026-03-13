# Docker-devops
#Project 2 - REST API - Python Application Code (Flask API)

python-docker-app
=======
**Python web application (portfolio + job apply links) and  Dockerize it.**
Technologies:
1. Python + Flask (lightweight web framework)
2. HTML + CSS
3. Buttons that redirect to Naukri.com and LinkedIn job search pages.

**Requiremnts :**
Python backend
Static frontend
Requirements file
Simple server

**Project Structure:**
asitav-portfolio/
│
├── app.py
├── requirements.txt
├── Dockerfile
│
├── templates/
│     └── index.html
│
└── static/
      └── style.css

Steps :
Step 1. Create Ec2 instance and connect from CLI
Step 2. sudo apt update -y
Step 3. install docker -> sudo apt install docker.io -y
note : Check the docker is activite or not --> sudo systemctl  status docker
step 4. docker run "Hello-world" --> Error permission denied because dockerd run in root user hence need to fix it 
step 5. sudo usermod -aG docker ubuntu(user)/ec2-user(user)
Explain : to add a user to the Docker group so they can run Docker commands without using sudo.
          sudo means Super User DO , usermod -usermod means modify a user account.
          -a means append -> to ensures the user is added to the group without removing existing groups.
          -G means Group -> It specifies which group you want to add the user to
step 6. in order to reflect the changes restart the docker means logout and log in 
step 7. create python basic program and save in app.py
step 8.create docker file 
step 8.1. Clone the Give repo
step 9. docker build -t asitavawsdevops/first-docker-project-image:latest .  (Build image)
Note: . --> means docker file in the same directory else need to specify 
docker run asitavawsdevops/python-application-image:latest -f dockerfile
step 10. check images --> docker images
step 11. docker run -it asitavawsdevops/first-docker-project-image:latest(run Container)
step 12  docker login 
          username: dockerhub username
          passwrod : dockerhub passwd
          docker run -d -p 5000:5000 asitavawsdevops/second-docker-project:latest (Push to docker Registory)
          
          Note: without -d -p 5000:5000  if we will run it will  fail because we need to map the port with the host port(ec2)
