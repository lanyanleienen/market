FROM alpine
ADD market-service /market-service
ENTRYPOINT [ "/market-service" ]
