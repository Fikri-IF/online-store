# My Online Shop

By : Fikri Ihsan Fadhiilah

## Installation
Make sure you installed docker

Install my-online-shop with

```bash
docker compose up
```


## API Reference
### User
#### Register

```http
  GET /api/v1/user/register
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required** username of account |
| `password` | `string` | **Required** password of account |

```http
{
    "status_code": 201,
    "message": "create new user successfully",
    "data": null
}
```

#### Login

```http
  POST /api/v1/user/login
```

| Request Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username` | `string` | **Required** username of account |
| `password` | `string` | **Required** password of account |

```http
{
    "status_code": 200,
    "message": "login successfull",
    "data": {
        "token": "eyJhbGciO1dadsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI2MzgxMTAsImlkIdasdaCJ1c2VybmFtZSI6InF3YWNrIn0.dTnOYC1fsT85xcmqasdwd2feCCxMbFCcyVIHAQ"
    }
}
```
### Category
#### Add new category
```http
  POST /api/v1/category
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required** category name |

```http
{
    "status_code": 201,
    "message": "category successfully created",
    "data": {
        "CategoryId": 12,
        "Name": "Beverages"
    }
}
```
#### Get all category
```http
  GET /api/v1/category
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "categories successfully fetched",
    "categories": [
        {
            "category_id": 2,
            "category_name": "Beverage"
        },
        {
            "category_id": 6,
            "category_name": "Fashion"
        }
    ]
}
```
#### Delete category
```http
  DELETE /api/v1/category/:categoryId
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "category successfully deleted",
    "data": null
}
```
### Product
#### Get product by id
```http
  GET /api/v1/product/:productId
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "product found",
    "data": {
        "product_id": 25,
        "product_name": "Bread",
        "price": 25000,
        "category": "Food",
        "stock": 100
    }
}
```
#### Get all product
```http
  GET /api/v1/product/
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "products successfully fetched",
    "products": [
        {
            "product_id": 29,
            "product_name": "Water Bottle",
            "price": 10000,
            "category": "Beverage",
            "stock": 500
        },
        {
            "product_id": 30,
            "product_name": "Cola",
            "price": 12500,
            "category": "Beverage",
            "stock": 300
        }
    ]
}
```
#### Get all product by category
```http
  GET /api/v1/product/category/:categoryId
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "products successfully fetched",
    "products": [
        {
            "product_id": 25,
            "product_name": "Bread",
            "price": 25000,
            "category": "Food",
            "stock": 100
        },
        {
            "product_id": 26,
            "product_name": "Pasta",
            "price": 15000,
            "category": "Food",
            "stock": 200
        },
        {
            "product_id": 27,
            "product_name": "Cheese",
            "price": 50000,
            "category": "Food",
            "stock": 150
        },
        {
            "product_id": 28,
            "product_name": "Apples",
            "price": 30000,
            "category": "Food",
            "stock": 120
        }
    ]
}
```
#### Add product
```http
  POST /api/v1/product/
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `product_name` | `string` | **Required** product name |
| `price` | `int` | **Required** product price |
| `category` | `int` | **Required** product category |
| `stock` | `int` | **Required** product stock |

```http
{
    "status_code": 201,
    "message": "product successfully created",
    "data": {
        "ProductId": 52,
        "ProductName": "apple",
        "Price": 2000,
        "Category": 2,
        "Stock": 2
    }
}
```
#### Delete product
```http
  DELETE /api/v1/product/:productId
```

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "product successfully deleted",
    "data": null
}
```
### Cart
#### Add to user cart
```http
  POST /api/v1/cart/
```
| Header | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required** token from login |

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `product_id` | `int` | **Required** product id |
| `quantity` | `int` | **Default 1** product quantity |

```http
{
    "status_code": 201,
    "message": "item successfully added",
    "data": null
}
```
#### Get user cart
```http
  GET /api/v1/cart/
```
| Header | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required** token from login |

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "cart items successfully fetched",
    "data": [
        {
            "product_id": 51,
            "product_name": "apple",
            "price": 2000,
            "quantity": 1
        }
    ]
}
```
#### Delete user cart
```http
  DELETE /api/v1/cart/
```
| Header | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required** token from login |

| Request Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `` | `` |  |

```http
{
    "status_code": 200,
    "message": "item deleted",
    "data": null
}
```
    
