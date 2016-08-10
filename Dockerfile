FROM alpine:3.3

COPY .dist/tugbot /usr/bin/tugbot

LABEL tugbot=leader

ENTRYPOINT ["/usr/bin/tugbot-leader"]