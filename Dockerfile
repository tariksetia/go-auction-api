################### Builder Stage  ######################################### 
FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN ls -l
RUN cd auction/api && go build -o auctionServer


#################### Final Stage  ##########################################
FROM alpine

WORKDIR /app
COPY --from=builder /go/src/auction/api/auctionServer .
ENTRYPOINT ./auctionServer
EXPOSE 8000
