# gRPC-Coffee

A simple gRPC demo project in Go that simulates a coffee shop — browse a menu, place an order, and check its status.

---

## Project Structure

```
gRPCCoffee/
├── client/         # gRPC client
├── server/         # gRPC server
├── proto/          # Protobuf definitions
├── gen/            # Generated Go protobuf code
├── go.mod
├── Makefile
└── README.md
```

---

## Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [protoc](https://grpc.io/docs/protoc-installation/) — Protocol Buffer compiler
- [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go)
- [protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc)

Install the protoc plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## Setup

```bash
git clone https://github.com/unsubstantiated-Script/gRPCCoffee.git
cd gRPCCoffee
go mod tidy
```

Regenerate protobuf code (if needed):

```bash
make proto
```

---

## Running

**Start the server** (listens on `:9001`):

```bash
go run ./server
```

**Run the client** (in a separate terminal):

```bash
go run ./client
```

Expected client output:

```
Menu: [id:"1" name:"Black Coffee"  id:"2" name:"Americano"  id:"3" name:"Vanilla Soy Chai Latte"]
Receipt: id:"ABC123"
Status: orderId:"ABC123"  status:"In Progress"
```

---

## RPC Reference

Defined in `proto/coffee_shop.proto`.

| RPC              | Request       | Response      | Description                        |
|------------------|---------------|---------------|------------------------------------|
| `GetMenu`        | `MenuRequest` | `Menu`        | Returns the list of available items |
| `PlaceOrder`     | `Order`       | `Receipt`     | Places an order, returns a receipt  |
| `GetOrderStatus` | `Receipt`     | `OrderStatus` | Returns the current order status    |

---

## Dependencies

| Package                          | Version   |
|----------------------------------|-----------|
| `google.golang.org/grpc`         | v1.79.3   |
| `google.golang.org/protobuf`     | v1.36.11  |
