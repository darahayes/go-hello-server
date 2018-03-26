FROM centos:7.4.1708
ARG BINARY=./go-hello-server
EXPOSE 4000

COPY ${BINARY} /opt/go-hello-server
CMD "/opt/go-hello-server"