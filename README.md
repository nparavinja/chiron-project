# chiron-project

building a docker image
 docker build --tag node-docker .

running a docker image
  docker run -d -p 8000:8000 node-docker


# front-end 
 docker build --tag angular-test .
 docker run -d -p 8000:80 angular-test

 http://localhost:4200

 TO DO: 
 - node_modules as volume, pretty slow build times
 - is building an image always needed (for any code change?)
 - start working on angular app
 - send api request from front-end