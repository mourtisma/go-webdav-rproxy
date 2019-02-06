FROM i386/golang:latest

RUN mkdir /app 
ADD /src /app/ 
WORKDIR /app

RUN go build

EXPOSE ${PORT}

CMD [ "./app" ]