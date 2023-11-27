# User Service

## USER_SERVICE ENDPOINT

### Registration

- users/register/user [POST]

- users/register/driver [POST]

- users/register/admin [POST]

### Log in

- users/login [POST]

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


## ORDER_SERVICE ENDPOINT

-