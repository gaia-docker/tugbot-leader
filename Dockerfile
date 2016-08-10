FROM alpine:3.3

COPY .dist/tugbot-leader /usr/bin/tugbot-leader

LABEL tugbot=leader

ENTRYPOINT ["/usr/bin/tugbot-leader"]