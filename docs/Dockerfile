FROM golang:1.10 as docs
RUN wget https://github.com/gohugoio/hugo/releases/download/v0.37.1/hugo_0.37.1_Linux-64bit.tar.gz && \
    tar -xzvf hugo_0.37.1_Linux-64bit.tar.gz && \
    mv hugo /usr/local/bin
WORKDIR /doc
COPY . /doc
RUN hugo

CMD ["/usr/local/bin/hugo", "server", "-D", "--bind", "0.0.0.0"]