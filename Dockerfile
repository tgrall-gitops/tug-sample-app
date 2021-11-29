FROM golang:1.17.2 as builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app/

RUN go build -o my-app

FROM scratch
ENTRYPOINT [ "/my-app" ]
COPY --from=builder /app/my-app /
COPY static /static/
