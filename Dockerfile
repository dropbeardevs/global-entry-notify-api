FROM alpine:3.16.2

WORKDIR /usr/local/bin

COPY ./bin/ .

CMD ["global-entry-notify-api"]