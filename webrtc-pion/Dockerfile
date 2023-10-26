FROM --platform=$TARGETPLATFORM golang:1.21.3-alpine
ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR /webrtc-pion
COPY go.mod go.sum ./
RUN go mod download
