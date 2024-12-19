FROM golang:1.23-alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm git ca-certificates

COPY . ./
RUN make install
RUN make build-prod

FROM scratch
COPY --from=builder /app/bin/prods /prods
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./prods" ]
