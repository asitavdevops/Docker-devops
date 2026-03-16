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

**Why We user Bind Mount/ Volume in docker :**

In Docker, containers are ephemeral (temporary). When a container is deleted, its data is also deleted.
To solve this, Docker provides Volumes and Bind Mounts so that data can persist outside the container.
Realtime Example of Volume requirement in  Real Production Example:
**A microservice architecture:**

Web Container
API Container
Database Container
        |
     Docker Volume
        |
Persistent Database Storage

Why volumes are used:
- Persistent data
- Data backup possible
- Containers can be recreated safely
Common use cases:
Databases (MySQL, PostgreSQL, MongoDB)
Logs
File storage
Shared application data

**===============Docker Bind Mount==========================**
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

**===============Docker Volume==========================**
A Docker Volume is managed by Docker itself, not directly by the host filesystem.
Diagram:
Host Machine
     |
     | /var/lib/docker/volumes/myvolume
     |
     └───────────────┐
                     │
              Docker Container
                     │
                  /app/data
                  
Docker stores volumes inside:/var/lib/docker/volumes/
Create Volume -> docker volume create myvolume
**Run Container with Volume :**
docker run -d \
  -v myvolume:/app/data \
  nginx

step 1. Volume Creation -> It can be created two ways 
1. Explicitly --> docker volume create myvolume
2. Automatic --> docker run -v myvolume:/data nginx (Docker automatically creates the volume if it does not exist)

step 2. Volume Attachment : A volume is attached when the container starts.
**docker run -v myvolume:/data nginx**

step 3.Data Persistence : Even if the container stops or is deleted, the data remains.
**docker rm container1**

step 4. Volume Sharing : Multiple containers can share the same volume.
**docker run -d -v myvolume:/data nginx**
**docker run -d -v myvolume:/data ubuntu**

step 5 . Volume Removal: Volumes are not deleted automatically unless explicitly removed.
**docker volume rm myvolume**

step 6. Remove Unused Volumes 
**docker volume prune**

**Volume Lifecycle Diagram:**

Why use volume? Example Because database data must persist even if the container crashes.

**Note: Docker Volue we can create in both Host OS and external Host like EC2,S3 etc**

**Practical Commands:**
1. Checking the running Volume --> docker volume ls
2. Create docker logical partition --> docker volume create asitav
3. How to see all details about volume :
  ** docker volume inspect asitav**
4. Delete a volume --> docker volume rm asitav
5. check top 5 docker images  --> docker images | head -5
6. create a docker image  and try to create a volume and mount 
   docker build . && docker run -it
   docker run -d --mount source=asitav , target=/app nginx:latest
   docker ps
   docker inspect <container-id>  (Here we trying inspect the entire container not the volume)
   Note : If the volume is used by any container then we cant directly delete/rm the volume
          First we have to stop the container OR delete the container then only we can able to delete the volume.

   **===========Docker Networking==================**
   Topics: Bridge vs Host vs Overlay |Secure containers with custom bridge network
   
   What is Docker Networking?
   Docker networking allows:
             1.Container → Container communication
             2.Container → Host communication
             3.Container → Internet communication
   Example: User → Nginx Container → Backend Container → Database Container
             Here Docker creates a network layer to connect them.

   **Bridge Network:**
   It allows you to bind a directory inside your container with the directory in your host operating system.
   So any file we are writing inside /app/folder(container) it will access/sync by the Host OS /app/Folder2 due to the bind mount.
   The Bydefault network in docker is the **bridge network** . So when any container created by defaulted bridge network got created
   If any container goes down the information wont be lost .

    **Architecture:**
                        Host Machine
                          |
                 Docker Bridge Network
                        docker0
                    (172.17.0.1)
               _________|___________
              |         |           |
          vethA      vethB       vethC
            |          |           |
        eth0@if10   eth0@if11   eth0@if12
            |          |           |
        Container1  Container2  Container3
        172.17.0.2  172.17.0.3  172.17.0.4

**To Check the Available Docker Networks : docker network ls**

  **Example Workflow:**
  Step 1.    Create 3  conatiners with name Login,Logout, FInance and Start those containers:
             docker run -dit --name Login nginx
             docker run -dit --name Logout nginx
             docker run -dit --name Finance nginx
Step 2.      Log in to the container and try to ping from Host to the container it should be successfull.
             -> docker ps 
             -> docker inspect Login (Gather the IP of log in container)
             -> docker network ls
             [-> docker network rm test ]
              connect to the conatner C1 to check if  Containers communicate using IP address.
             -> docker exec -it c1 ping 172.17.0.3
             ->ping from Host to C1_IP

             Log in to the container C1 and try to ping from C1 to the container C2 it should be successfull.
             -> docker ps 
             -> docker inspect Login (Gather the IP of log in container)
             -> docker inspect Logout (Gather the IP of log in container)
             -> docker network ls
             [-> docker network rm test ]
             ->ping from  to C1_IP
             NOte : Same Process for Finance conainer check 

**Custom Bridge Network (Secure & Recommended) :**
A custom bridge network improves security and container communication. When we want to create isolation with in the containers 

step 1 .Create Custom Bridge Network
          -> docker network create my_bridge
step 2. Check Network:
          -> docker network ls
step 3. Run Containers in Custom Network :
          -> docker run -dit --name web --network my_bridge nginx
          -> docker run -dit --name db --network my_bridge mysql
Note: Now containers communicate using container names.
Example: docker exec -it web ping db

**Why Custom Bridge is More Secure ?**
**Default bridge:**
Container A  ----\
Container B  ----- docker0 (all containers connected)
Container C  ----/

**Custom bridge:**
Network: frontend_net
   web1
   web2

Network: backend_net
   db1
   redis

**Here Containers in frontend_net cannot access backend_net unless explicitly connected.**
Ex - docker network connect backend_net web1

**Host Network**
In host networking, the container shares the host's network stack.
No separate IP.Container uses host IP directly.
**Architecture:**
Host Machine (192.168.1.10)
      |
   Container
      |
Uses same network as host

Ex: Run nginx with host network:
docker run --network host nginx
Access:
http://localhost:80

**Overlay Network (Docker Swarm / Multi Host)** :
Overlay networks connect containers across multiple Docker hosts.
Used in Docker Swarm or Kubernetes environments.

**Architecture:**
Host1                     Host2
 |                         |
Container A               Container B
      \                   /
       \----Overlay Network----

Containers communicate securely across hosts.

Example (Swarm):
**Initialize swarm:** 
-> docker swarm init
**Create overlay network:**
-> docker network create -d overlay my_overlay
**Run service:**
-> docker service create \
--name web \
--network my_overlay \
nginx

Note : Now containers on different machines communicate.

**What is Docker Swarm?**
Docker Swarm is Docker’s native container orchestration tool used to manage multiple Docker hosts as a single cluster.
**It helps with:**
Container orchestration
High availability
Load balancing
Scaling services
Rolling updates

Similar tools:Docker Swarm/Kubernetes (most popular today)

**Docker Swarm Architecture:**
                 Docker Swarm Cluster
                           |
                    Users / Internet
                           |
                      Load Balancer
                           |
                 ----------------------
                 |                    |
           Manager Node 1        Manager Node 2
           (Leader)              (Follower)
                 |
           Manager Node 3
           (Follower)

       -------------------------------------
       |             |             |       |
   Worker Node1  Worker Node2  Worker Node3 Worker Node4
       |             |             |         |
   Containers     Containers    Containers  Containers

   **Services in Docker Swarm :** A service defines how containers should run.
   Example: Swarm will run 3 containers automatically.
                    Service: Web App
                    Replicas: 3
                    Image: nginx
 ->  docker service create --name web --replicas 3 nginx
 Swarm ensures:
 web.1
 web.2
 web.3
If one container fails: 
web.2 ❌
Swarm automatically creates new container
web.4 ✔




   
    
   





















