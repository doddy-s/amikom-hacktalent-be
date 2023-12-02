FROM alpine:3.14

COPY amikom-hacktalent-be .

ENV MYSQL_HOST=127.0.0.1 MYSQL_USER=doddy-s MYSQL_PASSWORD=kain MYSQL_DBNAME=hacktalent

EXPOSE 3030

CMD ./amikom-hacktalent-be