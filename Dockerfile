# STEP 1 Prepare Go build  https://gist.github.com/Xeoncross/862d659e48b05d684dddad0ceaf36029
############################
#### Cachable. Using cached go mod download to speed up Docker builds https://petomalina.medium.com/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888
FROM golang:1.18beta1-alpine3.15 AS builder
WORKDIR /build
RUN apk add git
COPY go.mod .
COPY go.sum .
RUN go mod download

# STEP 2 Add the source code & build a static binary
# Non-cachable ###########
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /build/service

# STEP 3 Build image. Nobody user & SSH Certificates already build into the custom scratch image
###########################
FROM gcr.io/future-309012/scratch:latest
COPY --from=builder /build/service /service
EXPOSE 9090
USER nobody:nobody
ENTRYPOINT ["/service"]