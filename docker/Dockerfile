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
	drill

EXPOSE 9100
ADD api /
CMD ["./api"]
