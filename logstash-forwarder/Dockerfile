FROM digitalwonderland/base

RUN mkdir -p /etc/logstash \
 && curl -Lo /usr/local/bin/docker-logstash-forwarder https://github.com/digital-wonderland/docker-logstash-forwarder/releases/download/master/linux_amd64_docker-logstash-forwarder \
 && curl -Lo /usr/local/bin/logstash-forwarder https://github.com/digital-wonderland/docker-logstash-forwarder/releases/download/master/linux_amd64_logstash-forwarder \
 && chmod 0755 /usr/local/bin/{docker-logstash-forwarder,logstash-forwarder}

ADD logstash-forwarder.conf /etc/logstash/logstash-forwarder.conf 
RUN mkdir -p /mnt/logstash-forwarder
ADD logstash-forwarder.crt /mnt/logstash-forwarder/logstash-forwarder.crt
ADD logstash-forwarder.key /mnt/logstash-forwarder/logstash-forwarder.key

ENTRYPOINT ["/usr/local/bin/docker-logstash-forwarder"]
