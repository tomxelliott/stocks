# FROM golang:latest
# RUN mkdir /stocks
# ADD . /stocks/
# WORKDIR /stocks
# RUN go build -o main .
# CMD ["/stocks/main"]

FROM scratch
ADD main /
CMD ["/main"]