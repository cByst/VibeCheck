FROM golang:alpine as builder

WORKDIR /vibecheck

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /vibecheck/main

RUN echo "vibecheck:x:9090:9090:vibecheck:/var/empty:/sbin/nologin" > ./passwd

RUN echo "vibecheck:x:9090:vibecheck" > ./group

FROM scratch

COPY --from=builder /vibecheck/passwd /etc/passwd

COPY --from=builder /vibecheck/group /etc/group

COPY --from=builder /vibecheck/main /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

USER vibecheck

ENTRYPOINT ["/main"]