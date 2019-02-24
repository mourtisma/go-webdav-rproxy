ARG PROXY_IMAGE
FROM ${PROXY_IMAGE}

RUN mkdir /app 
ADD /src /app/ 
WORKDIR /app

RUN go build

ARG PORT
EXPOSE ${PORT}

CMD [ "./app" ]