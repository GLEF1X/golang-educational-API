package orm

import (
	"github.com/GLEF1X/golang-educational-API/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type ORM struct {
	dsn string
}

func (o *ORM) EstablishConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(o.dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("can't establish connection with database")
	}
	return db, nil
}

func (o *ORM) Migrate(conn *gorm.DB) {
	migrationErr := conn.AutoMigrate(&domain.Customer{})
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}
}

func NewORM(dsn string) *ORM {
	return &ORM{dsn: dsn}
}

func PrepareDatabase(dsn string) *gorm.DB {
	orm := NewORM(dsn)
	conn, err := orm.EstablishConnection()
	orm.Migrate(conn)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
