FROM centos
RUN yum install dlv -y \
    && yum install binutils -y \
    && yum install vim -y \
    && yum install gdb -y \
    && yum install git -y \
    && yum install wget -y

RUN wget https://dl.google.com/go/go1.14.12.linux-amd64.tar.gz  \
    && tar -C /usr/local -zxvf go1.14.12.linux-amd64.tar.gz \
    && rm /go1.14.12.linux-amd64.tar.gz

ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH /home
ENV GOROOT /usr/local/go
ENV GOPROXY goproxy.cn