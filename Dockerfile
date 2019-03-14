################### Builder Stage  ######################################### 
FROM golang:latest AS builder
WORKDIR /go/src/auction
COPY . .
RUN ls -l 
RUN cd ./api && go get -d ../...
RUN cd ./api && go build -o auctionServer


#################### Final Stage  ##########################################
FROM alpine

WORKDIR /app
COPY --from=builder /go/src/auction/api/auctionServer .
ENTRYPOINT ./auctionServer
EXPOSE 8000
