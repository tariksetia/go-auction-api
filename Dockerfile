################### Builder Stage  ######################################### 
FROM golang:latest AS builder
WORKDIR /go/src/
RUN git clone https://github.com/tariksetia/go-auction-api.git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN ls -l ./go-auction-api

WORKDIR /go/src/auction
RUN cp -r /go/src/go-auction-api/* .
RUN ls -l
RUN dep ensure
RUN cd ./api && go build -o auctionServer
RUN ls -l ./api

#################### Final Stage  ##########################################
FROM alpine

WORKDIR /app
COPY --from=builder /go/src/auction/api/auctionServer .
COPY --from=builder /go/src/auction/api/static/index.html .
RUN ls -l
ENTRYPOINT ./auctionServer
EXPOSE 8000
