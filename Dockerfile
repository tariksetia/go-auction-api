################### Builder Stage  ######################################### 
FROM golang:latest AS builder
WORKDIR /go/src/auction
COPY . .
RUN ls -l 
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN cd ./auction/api && go build -o auctionServer
RUN ls -l ./auction/api


#################### Final Stage  ##########################################
FROM alpine

WORKDIR /app
COPY --from=builder /go/src/auction/api/auctionServer .
COPY --from=builder /go/src/auction/api/index.html .
RUN ls -l
ENTRYPOINT ./auctionServer
EXPOSE 8000
