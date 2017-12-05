FROM golang:1.9.2-stretch
COPY echo.go /opt/
COPY build.sh /opt/
WORKDIR /opt
RUN go get github.com/golang/glog
RUN /opt/build.sh
