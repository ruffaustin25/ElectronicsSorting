docker build -t house-management .
docker run -it -p 1234:1234 --rm --name run-house-management house-management