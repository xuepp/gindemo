# ========================
# 构建阶段
# ========================
FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# 安装 swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init

# 静态编译 Go 程序
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# ========================
# 运行阶段
# ========================
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 6081

CMD ["./app"]
