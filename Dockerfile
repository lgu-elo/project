FROM alpine:latest

WORKDIR /app
COPY . /app

RUN apk add gcompat

EXPOSE 8083

ENTRYPOINT [ "/app/bin/project" ]
