FROM alpine:latest

EXPOSE 8080

FROM datadog/docker-dd-agent:11.3.585-dogstatsd-alpine

USER root

RUN apk update && apk add ca-certificates

USER 1001

RUN mkdir $DD_HOME/service

COPY ./auth $DD_HOME/service/
COPY ./supervisor.conf $DD_HOME/service/supervisor.conf
COPY ./sample_key.priv $DD_HOME/service/
COPY ./sample_key.pub $DD_HOME/service/

WORKDIR $DD_HOME/service

RUN cat $DD_HOME/service/supervisor.conf >> $DD_HOME/agent/supervisor.conf

CMD supervisord -n -c $DD_HOME/agent/supervisor.conf
