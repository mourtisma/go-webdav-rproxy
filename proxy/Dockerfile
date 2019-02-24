ARG PROXY_IMAGE
FROM ${PROXY_IMAGE}

RUN mkdir /app 
ADD /src /app/ 
WORKDIR /app

RUN go build

EXPOSE ${PORT}

CMD [ "./app" ]