ARG PROXY_IMAGE
FROM ${PROXY_IMAGE}

RUN mkdir -p /go/src 
ADD /src /go/src 
WORKDIR /go/src

RUN go get golang.org/x/oauth2/...

RUN go build -o app

ARG PORT
EXPOSE ${PORT}

CMD [ "./app" ]