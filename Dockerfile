FROM golang:1.16

WORKDIR /app

COPY ./ /app

RUN go build -o usermgmt-server ./usermgmt-server
RUN go build -o usermgmt-client ./usermgmt-client

EXPOSE 50051

CMD [ "./usermgmt-server" ] & [ "./usermgmt-client" ]
