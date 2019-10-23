FROM golang:1.12

ADD main.go /web/src/hello-world/
ADD glide.yaml /web/src/hello-world/
WORKDIR /web

RUN mkdir bin
ENV GOPATH /web
ENV PATH ${GOPATH}/bin:${PATH}

RUN curl -sL -o install_glide.sh https://glide.sh/get
RUN sh install_glide.sh

WORKDIR src/hello-world
RUN glide install
WORKDIR ..
RUN go install hello-world
CMD hello-world