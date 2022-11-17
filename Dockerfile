FROM alpine
ADD microuser-service /microuser-service
ENTRYPOINT [ "/microuser-service" ]
