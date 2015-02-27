docker build .

docker run -d -v /var/lib/docker:/var/lib/docker:ro -v /var/run/docker.sock:/var/run/docker.sock 254d0d178c97 -logstash="ec2-54-174-166-239.compute-1.amazonaws.com:5043"
