package data_store

type DataStore interface {
	GetUserNameById(userId string) (string, bool)
}
