FROM golang AS build
ENV location /Users/taeksoo/Development/practice/gRPCPractice/productinfo/src
WORKDIR ${location}/server

ENV GOARCH arm
ENV GOARM 7

ADD ./server ${location}/server
ADD ./proto ${location}/proto
ADD ./go.mod ${location}/go.mod
ADD ./go.sum ${location}/go.sum
#RUN go mod download

#RUN go get -d ./...
#RUN go install ./...
#RUN mkdir -p ~/go/src/golang.org/x
#RUN cd ~/go/src/golang.org/x
#RUN git clone https://github.com/golang/text.git 
#RUN rm -rf text/.git

RUN CGO_ENABLED=0 GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go build -o /bin/grpc-productinfo-server

# Build stage II
FROM scratch
COPY --from=build /bin/grpc-productinfo-server /bin/grpc-productinfo-server

ENTRYPOINT ["/bin/grpc-productinfo-server"]
EXPOSE 50051
