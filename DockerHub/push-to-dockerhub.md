# Log into Dockerhub
docker login --username=<dockerhubusername> (you will be prompted for your password)

# Tag the Docker image
docker tag gowebapi lsramanjaneyulu/gowebapi:latest

# Push the Docker image to DockerHub
docker push lsramanjaneyulu/gowebapi:latest