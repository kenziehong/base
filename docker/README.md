# 19. Starting a nginx web server 
  docker container --help
  docker container run --publish 80:80 --name webhost --detach nginx
  docker container logs webhost
  docker container top webhost
  docker container ls -a
  docker container stop 291
  docker container rm -f webhost

# 20. What happens when we run a container

# 21. Container vs VM. It’s just a process
https://www.bretfisher.com/getting-a-shell-in-the-docker-for-windows-vm/

  docker container run --name mongo -d mongo
  docker container ps
  docker run -it --rm --privileged --pid=host justincormack/nsenter1
  ps aux | grep mongo

# 23. Manage multiple containers
https://docs.docker.com/
--help

  docker container run -d --name mysql -e MYSQL_RANDOM_ROOT_PASSWORD=yes -p 3306:3306 mysql
  docker container run -d --name webserver -p 8080:80 httpd
  docker container run -d --name proxy -p 80:80 nginx
  docker container logs mysql
  docker container stop a831549061f4 1c8854c3c0eb 07769c784c8c 03e19554c569
  docker container rm a831549061f4 1c8854c3c0eb 07769c784c8c 03e19554c569
  docker container ls -a

# 25. What’s going on in containers. CLI process monitoring

  docker container inspect mysql (show config)
  docker container top mysql (list ps)
  docker container stats (monitor ps)

# 26. Getting a shell inside containers. No need for ssh
docker container run -it --name proxy nginx bash
ls -la
exit

docker container run -it --name ubuntu ubuntu
docker container start -ai ubuntu (rerun)
docker container exec -it mysql bash (run additional process in running container)
docker pull alpine (distribution linux, with small size, 5M, apk package manager)
docker container run -it alpine sh

apt-get update (apt-get package manager in ubuntu)
apt-get install -y curl
curl google.com
apt-get install -y procps
ps aux
exit

# 27. Docker networks: Concepts for private and public commns in containers
docker container run -d --name webhost -p 80:80 nginx
docker container port webhost
docker container inspect --format "{{.NetworkSettings.IPAddress}}" webhost

# 29. Docker networks: CLI management of virtual 
docker network ls
docker network inspect bridge
docker network create my_app_net
docker network inspect my_app_net
docker container start webhost
docker container run -d --name new_nginx --network my_app_net nginx
docker network connect 20f9d1802348 b406cc3d7980
docker network disconnect 20f9d1802348 b406cc3d7980

# 30. Docker Networks: DNS and how containers find each other

# 32. Assignment Answers: Using Containers for CLI Testing
docker container run --rm -it ubuntu:14.04 bash
apt-get upate && apt-get install -y curl
docker container run --rm -it centos:7 bash
yum update curl
curl --version

// ANCHOR 34. Assignment: DNS round robin test
# 34. Assignment: DNS round robin test
docker network create dude
docker container run -d --net dude --net-alias search elastisearch:2
docker container run -d --net dude --net-alias search elasticsearch:2
docker container ls
docker container run --rm --net dude alpine nslookup search
docker container run --rm --net dude centos curl -s search:9200

// ANCHOR  36. What’s in an image
# 36. What’s in an image

# 37. Using docker hub registry images

# 38. Discover the image cache

# 39. Image tagging and pushing image

# 40. Building image 


