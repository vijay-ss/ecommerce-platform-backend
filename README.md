# ecommerce-platform-backend
Backend for an Ecommerce platform written in Go.

## Architecture
![](img/architecture.png)

## Tools
- Go
- Gin Framework
- JWT Authentication
- Bcrypt for password hashing
- MongoDB
- Docker
- Postman

## API Endpoints

- http://localhost:8000/users/signup
- http://localhost:8000/users/login
- http://localhost:8000/admin/addproduct
- http://localhost:8000/admin/productview
- http://localhost:8000/users/search

Requires Authentication:
- http://localhost:8000/addtocart
- http://localhost:8000/removeitem
- http://localhost:8000/listcart
- http://localhost:8000/addaddress
- http://localhost:8000/edithomeaddress
- http://localhost:8000/editworkaddress
- http://localhost:8000/deleteaddress
- http://localhost:8000/cartcheckout
- http://localhost:8000/instantbuy

## Testing a few Endpoints

You can spin up the project locally with the following commands:

```bash
docker-compose up
go run main.go
```

- **User signup (POST REQUEST)**

http://localhost:8000/users/signup

```json
{
    "first_name": "Iron",
    "last_name": "Mike",
    "email": "ironmike@gmail.com",
    "password": "ironmike123",
    "phone": "123456789"
}
```
Response :"Successfully signed up!"

- **User login (POST REQUEST)**

  http://localhost:8000/users/login

```json
{
  "email": "ironmike@gmail.com",
  "password": "ironmike123"
}
```

The response will have the following format:
```json
{
    "_id": "65d3ba7df35272fd17f4b4ec",
    "first_name": "Iron",
    "last_name": "Mike",
    "password": "$2a$14$/mXZr3Dqld4xeDIbOE391e9/DcDI4aoPql3S6Boxj7e6oei4GWd6e",
    "email": "ironmike@gmail.com",
    "phone": "123456789",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Imlyb25taWtlQGdtYWlsLmNvbSIsIkZpcnN0TmFtZSI6Iklyb24iLCJMYXN0TmFtZSI6Ik1pa2UiLCJVaWQiOiI2NWQzYmE3ZGYzNTI3MmZkMTdmNGI0ZWMiLCJleHAiOjE3MDg0NjEwNzF9.iL6h8Cj79Me0pKbp2a3_qEA5knxnoRtRZlr5xkna5TU",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0TmFtZSI6IiIsIkxhc3ROYW1lIjoiIiwiVWlkIjoiIiwiZXhwIjoxNzA4OTc5NDcxfQ.cWR6WYZJDOnT5FeIVDd6y0ewhNBuOFmXSHTfzrKBd9o",
    "created_at": "2024-02-19T20:30:53Z",
    "updated_at": "2024-02-19T20:30:53Z",
    "user_id": "65d3ba7df35272fd17f4b4ec",
    "usercart": [],
    "address": [],
    "orders": []
}
```

- **Admin add product to available products (POST REQUEST)**

http://localhost:8000/admin/addproduct

```json
{
  "product_name": "Macbook M3 Pro",
  "price": 3500,
  "rating": 10,
  "image": "macbook_pro_m3.jpg"
}
```
Response : "successfully added product"

- **View all available products GET REQUEST**

http://localhost:8000/users/productview

The response will have the following format:
```json
[
    {
        "Product_ID": "65d3baeaf35272fd17f4b4ee",
        "product_name": "Alienware x17",
        "price": 3000,
        "rating": 10,
        "image": "alienware.jpg"
    },
    {
        "Product_ID": "65d3bb18f35272fd17f4b4f0",
        "product_name": "iPhone 15 Pro Max",
        "price": 1500,
        "rating": 8,
        "image": "iphone15_pro_max.jpg"
    },
    {
        "Product_ID": "65d3bb3cf35272fd17f4b4f2",
        "product_name": "Pixel 8 Pro",
        "price": 1200,
        "rating": 9,
        "image": "pixel8_pro.jpg"
    },
    {
        "Product_ID": "65d3bb5df35272fd17f4b4f4",
        "product_name": "Macbook M3 Pro",
        "price": 3500,
        "rating": 10,
        "image": "macbook_pro_m3.jpg"
    }
]
```

- **Search for product(s) (GET REQUEST)**

Find products with a specific keyword
http://localhost:8000/users/search?name=Pro

The response will have the following format:
```json
[
    {
        "Product_ID": "65d3bb18f35272fd17f4b4f0",
        "product_name": "iPhone 15 Pro Max",
        "price": 1500,
        "rating": 8,
        "image": "iphone15_pro_max.jpg"
    },
    {
        "Product_ID": "65d3bb3cf35272fd17f4b4f2",
        "product_name": "Pixel 8 Pro",
        "price": 1200,
        "rating": 9,
        "image": "pixel8_pro.jpg"
    },
    {
        "Product_ID": "65d3bb5df35272fd17f4b4f4",
        "product_name": "Macbook M3 Pro",
        "price": 3500,
        "rating": 10,
        "image": "macbook_pro_m3.jpg"
    }
]
```

- **Adding product to cart (GET REQUEST)**

http://localhost:8000/addtocart?id=65d3bb5df35272fd17f4b4f4&userID=65d3ba7df35272fd17f4b4ec

- **Cart Checkout and placing the order (GET REQUEST)**

  After placing the order the items have to be deleted from cart functonality added

  http://localhost:8000?id=xxuser_idxxx

- **Instant buy of product(s) (GET REQUEST)**
  http://localhost:8000?userid=65d3ba7df35272fd17f4b4ec&pid=65d3bb5df35272fd17f4b4f4
