Getting Started
===
1. [Install Docker](https://docs.docker.com/get-started/)
2. Run TestBuildAndRun.sh on your machine
3. Go to localhost:1234 and make sure you see the webpage
4. Make your changes and push them
5. ssh into the server using "ssh -l root 192.168.0.2"
6. build and run
    1. Open the Docker tab in the Unraid server web portal and you should be able to see the "HouseManagment" application running, left click on the image on the left of that applicaiton and select "Stop"

    1. "sudo docker build -t house-management ."
            
        This command runs the docker build process on the current directory and names the output "house-management".  The docker build process is determined by the Dockerfile

    2. "sudo docker run -it -p 1234:1234 -d --rm --name HouseManagement house-management"

        This command runs the docker container named "house-management" in the background, calls the running instance "HouseManagement", opens the port 1234 and connects it to port 1234 on the containerized guest

    3. Go to 192.168.0.2:1234 to make sure that the web server is running