# Api Gateway

## User Service Endpoints

### Registration

- users/register/user [POST]

- users/register/driver [POST]

- users/register/admin [POST]

### Log in / out

- users/login [POST]

- users/logout [GET]

### Verification

- users/verify/:userId/:token

## Merchant Services Endpoints

### Create Restaurant and Menu

- merchant/restaurant [POST] (Requires auth with admin role)

- merchant/restaurant [PUT] (Requires auth with admin role)

- merchant/menu [GET] (Requires auth with admin role)

- merchant/menu/:id [GET] (Requires auth with admin role)

- merchant/menu [POST] (Requires auth with admin role)

- merchant/menu/:id [PUT] (Requires auth with admin role)

- merchant/menu/:id [DELETE] (Requires auth with admin role)

### View Restaurants and Menus as Customer

- restaurants [GET] 

- restaurants/:resto_id [GET] 

- restaurants/menu/:menu_id [GET]

## Order Service Endpoints

### Make & view order(s) as a Customer

- users/orders [POST] (Requires auth with user role)

- users/orders [GET] (Requires auth with user role)

- users/orders/:id [GET] (Requires auth with user role)

### As a Merchant

- merchant/orders [GET] (Requires auth with admin role)

- merchant/orders/ongoing [GET] (Requires auth with admin role)

- merchant/orders/:id [GET] (Requires auth with admin role)

- merchant/orders/:id [PUT] (Requires auth with admin role)

> [PUT] updates the order status such as (done, cancelled)

### As a Driver

- driver/orders [GET] (Requires auth with driver role)

- driver/orders/ongoing [GET] (Requires auth with driver role)

- driver/orders/:id [PUT] (Requires auth with driver role)

> [PUT] updates the order status such as (delivary, done)