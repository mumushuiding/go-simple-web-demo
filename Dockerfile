FROM scratch
ADD /go-simple-web-demo //
ADD /config.json //
EXPOSE 8080
ENTRYPOINT [ "/go-simple-web-demo" ]