from alpine:latest

RUN apk	add curl \
	iproute2 \
	iputils \
	bash bash-doc bash-completion \
	util-linux \
	pciutils \
	usbutils \
	coreutils \
	binutils \
	findutils \
	grep \
	tcpdump \
	mysql-client \
	postgresql-client \
	bind-tools \
	drill

EXPOSE 9100
CMD ["/bin/sh", "-c", "while true; do sleep 5; done;"]
