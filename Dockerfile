# build stage
FROM ubuntu:20.04 AS builder
RUN apt-get update && DEBIAN_FRONTEND=noninteractive && apt-get install -y --no-install-recommends tzdata
RUN apt-get -y install golang-go
RUN useradd -m worker
RUN mkdir /home/worker/models && \
    mkdir /home/worker/controllers && \
    chown -R worker:worker /home/worker
WORKDIR /home/worker/
USER worker
ADD main.go go.mod ./
ADD controllers/ ./controllers
ADD models/ ./models
RUN go build .

# final stage
FROM debian:buster-slim
ENV LC_ALL=C.UTF-8
RUN apt update && apt upgrade -y && \
    useradd -m worker
USER worker
WORKDIR /home/worker/
COPY --from=builder /home/worker/container-state-service .
USER worker
ENTRYPOINT ["./container-state-service"]