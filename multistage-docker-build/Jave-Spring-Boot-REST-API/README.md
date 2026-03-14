#Project Overview :
We will create a simple Spring Boot REST API.
Endpoint: http://localhost:8080
Output : Welcome to Asitav DevOps Java Application

Project Structure :
java-docker-app
│
├── src
│   └── main
│       └── java
│           └── com
│               └── example
│                   └── demo
│                       └── DemoApplication.java
│
├── pom.xml
└── Dockerfile

step 1. pom.xml (Maven Configuration)
        Create pom.xml

Check the pom.xml it performs below behaviour 
This installs:
Spring Boot web server
Embedded Tomcat

step 2. Java Application Code :
Create file: src/main/java/com/example/demo/DemoApplication.java
and put the code .

step 3. Build the Application
        Install Maven if not installed.
        Run --> mvn clean package 
    This creates: target/java-docker-app-1.0.jar 
    This JAR file will run the application.
Test locally: java -jar target/java-docker-app-1.0.jar
open : http://localhost:8080

step 4. Build the Application 
        RUN : mvn clean package
    This creates: target/java-docker-app-1.0.jar
    This JAR file will run the application.

step 5 . create multi stage docker file and  containerize it 