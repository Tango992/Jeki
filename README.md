# Jeki

[Feel free to try this API. Click to access the documentation.](https://jeki-egmflbdzpa-et.a.run.app)

## Description: 

Food delivery application that integrates customer, driver, and restaurant services seamlessly.

## Background:

> Coming from curiosity, "What does it takes to develop an application like GoFood / Grab Food", we tried to reverse engineer the infrastructure and develop our own food delivery application using Go. Utilizing industry standards tech stacks, our small team with 2 weeks timeline, tons of exploration, we managed to execute this application successfully. We learned a lot about building this kind off application, and surely this is our most complex application yet that utilizes microservices.

## Highlights:

* Microservices Architecture
* Google Maps Integration
* Serverless Deployment with Google Cloud Run
* Payment Gateway (Xendit)
* Email notifications

### Tech stacks:

* Go
* Echo
* gRPC
* Docker
* PostgreSQL
* MongoDB
* Redis
* JWT-Authorization
* 3rd Party APIs (Xendit, Google Maps)
* SMTP
* REST
* Swagger
* Testify

## Application Flow

![Final Flow](./misc/flow.svg)

## ERD

### User Service (Postgres)

![User service ERD](./misc/user_erd.svg)

### Merchant Service (Postgres)

![ERD](./misc/merchant_erd.svg)

### Order Service (MongoDB)

![ERD](./misc/order_erd.svg)
