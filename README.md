# Project: Bookstore APIs
  This project is to support bookstore CRUD operations of shopping cart. 
  This is a demo project, it does not support operation related to order or product

# Local Development Notes
1. Need to download docker and postgres image for docker to initialize data using instruction in data folder
2. Once db is ready, import the postman_collections.json to run test on this api

# 📁 Collection: server url/api/cart 


## End-point: Create Cart
This endpoint allows you to create a cart or add items to a specific cart identified by the cart ID. Upon a successful execution, it returns the updated cart details including the cart ID and the list of items added to the cart.

The request should be a POST method to the endpoint `{{base_url}}/cart/8616ca33-b7aa-438e-bb45-862c6406ede9`.

The response will have a status code of 201 for creating a cart or status code 200 for adding items to the cart. The response body will contain the updated cart details, including the cart ID and the list of items added to the cart, where each item is represented by its SKU ID and the quantity added.

Example response body:

``` json
{
    "cartId": "5d2f8a5e-9379-4e74-8e6c-f8027eea28aa",
    "cartItems": [
        {
            "skuId": "2b8689ca-4c9e-417c-9303-f012ab75f525",
            "quantity": 2
        }
    ]
}

 ```
### Method: POST
>```
>{{base_url}}/cart?userId=5d2f8a5e-9379-4e74-8e6c-f8027eea28aa
>```
### Query Params

|Param|value|
|---|---|
|userId|5d2f8a5e-9379-4e74-8e6c-f8027eea28aa|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update Cart
This is a POST request, submitting data to an API via the request body. This request submits JSON data, and the data is reflected in the response.

A successful POST request typically returns a `200 OK` or `201 Created` response code.
### Method: POST
>```
>{{base_url}}/cart/5d2f8a5e-9379-4e74-8e6c-f8027eea28aa
>```
### Headers

|Content-Type|Value|
|---|---|
|userId|5d2f8a5e-9379-4e74-8e6c-f8027eea28aa|


### Body (**raw**)

```json
{   "cartId":"5d2f8a5e-9379-4e74-8e6c-f8027eea28aa",
    "cartItems": [
        {
            "skuId":"2b8689ca-4c9e-417c-9303-f012ab75f525",
            "quantity": 3
        },
                {
            "skuId":"eb5e507c-ff58-4750-be22-5f8dd52fb4ec",
            "quantity": 2
        }
    ],
    "total": 25.00
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Cart Items
This endpoint retrieves the details of a specific shopping cart identified by its ID. The response will include the cart ID and an array of cart items, each containing the SKU ID and the quantity.

The response will have a status code of 200, and the content type will be in JSON format.
### Method: GET
>```
>{{base_url}}/cart/5d2f8a5e-9379-4e74-8e6c-f8027eea28aa?userId=08d92108-b49d-4aa6-b3ee-d6248813bad4
>```
### Query Params

|Param|value|
|---|---|
|userId|08d92108-b49d-4aa6-b3ee-d6248813bad4|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Checkout Cart
### Method: POST
>```
>{{base_url}}/cart/5d2f8a5e-9379-4e74-8e6c-f8027eea28aa/checkout?userId=5d2f8a5e-9379-4e74-8e6c-f8027eea28aa
>```
### Query Params

|Param|value|
|---|---|
|userId|5d2f8a5e-9379-4e74-8e6c-f8027eea28aa|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
