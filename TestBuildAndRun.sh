docker build -t electronics-sorting .
docker run -it -p 2796:2796 --rm --name run-electronics-sorting electronics-sorting