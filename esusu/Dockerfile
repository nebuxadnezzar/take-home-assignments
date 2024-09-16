FROM alpine

COPY ./dist/static/* /app/
WORKDIR /app
CMD ["/app/maas", "8080", "./users_flushed.json"]