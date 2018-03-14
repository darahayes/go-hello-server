FROM centos:7.4.1708
ARG BINARY=./go-hello-server
EXPOSE 8000

COPY ${BINARY} /opt/go-hello-server
ENTRYPOINT ["/opt/go-hello-server"]