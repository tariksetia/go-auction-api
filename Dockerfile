################### Builder Stage  ######################################### 
FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN ls -l 
RUN cd ./auction/api && go get -d ../...
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
