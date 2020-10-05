FROM golang AS build
ENV location /Users/taeksoo/Development/practice/gRPCPractice/productinfo/src/ 

ENV GOARCH arm
ENV GOARM 7

WORKDIR ${location}/client

ADD ./client ${location}/client
ADD ./proto ${location}/proto
ADD ./go.mod ${location}/go.mod
ADD ./go.sum ${location}/go.sum
#RUN go mod download
#
#RUN go get -d ./...
#RUN go install ./...

RUN CGO_ENABLED=0 GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go build -o /bin/grpc-productinfo-client

# Build stage II
FROM scratch
COPY --from=build /bin/grpc-productinfo-client /bin/grpc-productinfo-client

ENTRYPOINT ["/bin/grpc-productinfo-client"]
#EXPOSE 50052
