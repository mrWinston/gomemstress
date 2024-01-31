FROM quay.io/app-sre/boilerplate:image-v4.0.2 AS builder

ENV OPERATOR_PATH=/go/src/github.com/mrWinston/memstress
ENV GO111MODULE=on
ENV GOFLAGS=""

COPY . ${OPERATOR_PATH}
WORKDIR ${OPERATOR_PATH}

RUN go build

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.9-1108.1705420507
ENV OPERATOR_BIN=memstress

WORKDIR /root/
COPY --from=builder /go/src/github.com/mrWinston/memstress/${OPERATOR_BIN} /usr/local/bin/${OPERATOR_BIN}
ENTRYPOINT /usr/local/bin/memstress
