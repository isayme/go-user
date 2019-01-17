FROM alpine

WORKDIR /opt/bin

COPY ./bin/user .

CMD ["./user"]
