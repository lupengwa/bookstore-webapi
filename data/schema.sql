CREATE SCHEMA bookstore;
CREATE TABLE bookstore.users (
   id UUID PRIMARY KEY,
   email VARCHAR(255) UNIQUE,
   created_at timestamp default current_timestamp,
   updated_at timestamp
);


CREATE TABLE bookstore.books (
      id UUID PRIMARY KEY,
      name VARCHAR UNIQUE,
      price NUMERIC (12,2),
      total INT,
      created_at timestamp default current_timestamp,
      updated_at timestamp
);

CREATE TABLE bookstore.orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    total NUMERIC (12,2) NOT NULL,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    FOREIGN KEY (USER_ID) references bookstore.users(ID)
);

CREATE TABLE bookstore.order_items (
    order_id UUID NOT NULL,
    sku_id  UUID NOT NULL,
    quantity SMALLINT NOT NULL,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    PRIMARY KEY (order_id, sku_id),
    FOREIGN KEY (SKU_ID) references bookstore.books(ID),
    FOREIGN KEY (ORDER_ID) references bookstore.orders(ID)
);

CREATE TABLE bookstore.cart(
   id UUID PRIMARY KEY,
   total NUMERIC (12,2) NOT NULL,
   created_at timestamp,
   updated_at timestamp,
   CONSTRAINT fk_user
       FOREIGN KEY (id)
           REFERENCES bookstore.users (id)
);


CREATE TABLE bookstore.cart_items(
   cart_id UUID NOT NULL,
   sku_id  UUID NOT NULL,
   quantity SMALLINT NOT NULL,
   created_at timestamp,
   updated_at timestamp,
   PRIMARY KEY (cart_id, sku_id),
   FOREIGN KEY (SKU_ID) references bookstore.books(ID),
   FOREIGN KEY (CART_ID) references bookstore.cart(ID)
);
