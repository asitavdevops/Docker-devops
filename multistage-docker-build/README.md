**Multi Staging Docker Concepts:**
When we create a normal docker file then due to the base image and dependencies and library build images takes huge disk size and more over 
out applcaiton may expose to OS Vulnerability . 
**Dockerfile OS Vulnerability means security vulnerabilities** that come from the base operating system used in a Docker image. 
If the base image has vulnerabilities, your container will also inherit them.
Example base images from the Docker ecosystem: Ubuntu ,Alpine Linux,Debian,Python

**Q> What are the Production issue that you faced with docker container and how did you solve it .**
**Ans -** Previously we are using ubuntu base images and even in the final stage we are using ubuntu or python runtime images which are 
exposed to some kind of issues with this images so we move to distroless images(Im my organization we moved to python distroless images
which only had python runtime and because of this we didnot have basic package like fild,ls, wget,curl etc) so it was providing us 
highlevel of security and after deploying the application we can say that our application is not exposed to any OS Vulnerability and
size of the image also drastically reduce upto 800%.

**Best Answer:** One production issue we faced with Docker containers was security vulnerabilities in the base images.
Initially, our containers were built using base images like Ubuntu or full Python images. These images contain many OS packages, 
which introduced multiple OS-level vulnerabilities (CVEs) and increased the overall container size.
To solve this problem, we migrated to Distroless images, specifically Python distroless images.
Distroless images contain only the application runtime and required dependencies, without unnecessary utilities like:

bash
curl
wget
ls
apt

Because of this:
1. The attack surface was significantly reduced
2. The number of OS vulnerabilities decreased
3. Security posture of our production containers improved
4. After migrating to distroless images, we also integrated container vulnerability scanning using Trivy
   in our CI/CD pipeline to ensure images are secure before deployment.

**Follow up question: But how do you debug distroless containers?**
Since distroless images don't include debugging tools, we usually debug using temporary debug containers or reproduce 
the issue locally using a similar base image that includes shell utilities.
Example of temporary debug containers
if something fails in production, you cannot log into the container to troubleshoot. Below command will fail because bash is not available.
--> docker exec -it my-container bash
Run a temporary debug container in the same network:
-->docker run -it --network container:my-container ubuntu bash
Now you can run commands like: 
</> Bash
curl localhost:8080
ping database


