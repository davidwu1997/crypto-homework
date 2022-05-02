FROM golang:1.16 as builder

WORKDIR /crypto
COPY . /crypto

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go build -mod=readonly -v -o server main.go

FROM alpine:3

RUN apk add --no-cache ca-certificates
RUN mkdir -p /crypto
WORKDIR /crypto

COPY --from=builder /crypto/server /crypto/server
COPY --from=builder /crypto/deployment /crypto/deployment

# Create appuser.
ENV USER=appuser
ENV UID=51197

RUN adduser \
--disabled-password \
--gecos "application user" \
--no-create-home \
--uid "${UID}" \
"${USER}"

RUN chown appuser:appuser /crypto
RUN chown appuser:appuser /crypto/*
USER appuser:appuser

ENTRYPOINT ["/crypto/server"]