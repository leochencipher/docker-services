FROM progrium/busybox
MAINTAINER philcryer phil@philcryer.com

ENV DOCKER_HOST unix:///tmp/docker.sock

ADD start    /bin/start
ADD beaver.conf /tmp/beaver.conf

RUN chmod 755 /bin/start && \
	opkg-install python && \
	opkg-install python-pip && \
	pip install beaver

ENTRYPOINT [ "/bin/sh" ]
CMD [ "/bin/start" ]
