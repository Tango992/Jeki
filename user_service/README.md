# User Service

## Endpoint

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

- restaurants [GET] (Requires auth with user role)

- restaurants/:resto_id [GET] (Requires auth with user role)

- restaurants/:resto_id/menu/:menu_id [GET] (Requires auth with user role)