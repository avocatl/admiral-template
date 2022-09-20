FROM golang:1-alpine as base

ENV CGO_ENABLED=0

WORKDIR /cli

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o cli cmd/cli/main.go

FROM base as testing
ENTRYPOINT [ "go", "test", "./...", "-v"]

FROM scratch as runner
COPY --from=base /cli/cli /
ENTRYPOINT [ "/cli" ]