FROM alpine:latest
ENV TZ="America/New_York"
ENV ZONEINFO="/zoneinfo.zip"
COPY ./app .
COPY ./zoneinfo.zip .
CMD ["/app"]