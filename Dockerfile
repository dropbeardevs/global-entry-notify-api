FROM golang:1.19

WORKDIR /usr/local/bin

COPY ./bin/ .

CMD ["global-entry-notify-api"]