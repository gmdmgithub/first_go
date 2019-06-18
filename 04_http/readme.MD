# Simple web apps but dokerized

## create htpp sample - exposed on some port (8081 in the sample)

## expose in docker

1. Create a Dockerfile

2. Run docker build - dot at the end is important

``` bush
$ docker build -t <NAME> .
```
3. Check if the image was created
``` bush
$ docker images
to remove image
$ docker rmi <NAME>  (if force necessary use '-f' flag)
```
4. Run the app in a docker (first in interactive mode)
``` bush
docker run -p 8080:8081 -it <NAME>
```
5. Run in detatched mode (background)
``` bush
docker run -p 8080:8081 -d <NAME>
```
6. Check docker - process should be in the background
``` bush
docker ps
```