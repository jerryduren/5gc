FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o ts29502nsmfpduserver .

FROM scratch AS runtime
COPY --from=build /go/src/ts29502nsmfpduserver ./
EXPOSE 8080/tcp
ENTRYPOINT ["./ts29502nsmfpduserver"]
