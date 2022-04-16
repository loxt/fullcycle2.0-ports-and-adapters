FROM golang:1.17

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0

RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1001 www-data

USER www-data

CMD ["tail", "-f", "/dev/null"]