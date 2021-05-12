FROM busybox:glibc

EXPOSE 8972

WORKDIR /
COPY ./hashpower /hashpower
ENTRYPOINT ["/hashpower"]
# CMD ["-h"]

