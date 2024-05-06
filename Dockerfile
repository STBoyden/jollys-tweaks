FROM golang:alpine as BUILDER
COPY . .
RUN go build -o /bin/app main.go

FROM alpine:latest as RUNNER
COPY --from=BUILDER /bin/app /bin/app
VOLUME out
CMD ["/bin/app", "-out=/out"]