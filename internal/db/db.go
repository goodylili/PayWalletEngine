package db

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {

}
