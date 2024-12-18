FROM golang:1.23-alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm git

COPY . ./
RUN make install
RUN make build-prod

FROM scratch
COPY --from=builder /app/bin/prods /prods

ENTRYPOINT [ "./prods" ]
