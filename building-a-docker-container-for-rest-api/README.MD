# HOW TO RUN LOCALLY
1. Init PORT, using "export PORT=:8080"
2. Check PORT, using "echo $PORT"
3. Check file building-a-docker-container-for-rest-api, if doesn't exist, run "go build"
4. After all, run "./building-a-docker-container-for-rest-api"
5. Now server will running locally

# HOW TO RUN USING DOCKER
1. Check docker images golang-api exist or not, using "docker images"
2. If not exist run "docker build -t golang-api ." and will build docker images 
3. After build completed, run "docker run -p 8080:8000 golang-api"
4. Now server will running using Docker