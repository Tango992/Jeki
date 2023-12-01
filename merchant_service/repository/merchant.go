package repository

type Merchant interface{
	FindMultipleMenuDetails() error
}