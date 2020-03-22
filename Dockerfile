FROM alpine:3.11

RUN apk add --no-cache ca-certificates

ADD ./our-life-before-corona /our-life-before-corona
ADD ./database_seed /database_seed

ENTRYPOINT ["/our-life-before-corona"]
