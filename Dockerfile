FROM alpine:latest

RUN adduser -S dockeruser
RUN mkdir /service
COPY ./auth /service/auth
COPY ./sample_key.priv /service/
COPY ./sample_key.pub /service/
RUN chown -R dockeruser /service

USER dockeruser
WORKDIR /service

CMD /service/auth
