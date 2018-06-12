FROM golang:onbuild

RUN mkdir /build
WORKDIR /build
COPY main.go /build
COPY bot.go /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rolebot .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN touch /rolebot.conf
WORKDIR /root
COPY --from=0 /build/rolebot .
CMD ["/root/rolebot", "-conf", "/rolebot.conf"]
