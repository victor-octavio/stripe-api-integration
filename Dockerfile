FROM golang:1.23-alpine as stage1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o apiExec

FROM gcr.io/distroless/static:nonroot

COPY --from=stage1 /app/apiExec /

ENTRYPOINT ["/apiExec"]