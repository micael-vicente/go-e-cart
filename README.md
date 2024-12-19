# Go e-cart

A very simple cart service written in Go.
The service is a REST API that allows you to create, list, update and delete carts and items.

## Requirements 📖
- Go 1.23.4 
- Docker
- Docker Compose

## Running ▶️

### Starting the service

```bash
docker-compose -f docker/docker-compose.yml up
```

```bash
go run cmd/cart/main.go
```

### Stopping the service

Send a `SIGINT` signal to the process.

```bash
docker-compose -f docker/docker-compose.yml down
```

## Usage

### Create a cart

```bash
curl -X POST http://localhost:8081/carts?customerId=CUST1
```

### List all carts

```bash
curl -X GET http://localhost:8081/carts?customerId=CUST1
```

### Get a cart by ID

```bash
curl -X GET http://localhost:8081/carts/1
```

### Update the items of a cart

```bash
curl -X PUT http://localhost:8081/carts/1 \
-d '[{"sku": "SKU1", "quantity": 2}, {"sku": "SKU2", "quantity": 1}]' \
-H "Content-Type: application/json"
```

### Delete a cart

```bash
curl -X DELETE http://localhost:8081/carts/1
```



