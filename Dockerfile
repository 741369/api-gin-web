FROM centos:7

WORKDIR /api-gin-web

ADD api-gin-web /api-gin-web/

CMD ["/api-gin-web/api-gin-web"]