// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package pgdb

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

type DBConfig struct {
	Addr             string
	User             string
	Password         string
	Database         string
	DBModel          []interface{}
	DBCompositeTypes []interface{}
}

type DBComponent struct {
	config *DBConfig
	db     *pg.DB
	prod   bool
}

func NewDBComponent(dbConfig *DBConfig, prod bool) *DBComponent {
	// CIRA = Construction = Initialization = Return (Resource) Allocation
	// 1. Construction
	dbComp := &DBComponent{
		config: dbConfig,
		prod:   prod,
	}
	// 2. Initialization & Connect
	dbComp.ConnectDataBase()
	if dbConfig.DBModel != nil { // if we don't have a schema, resume w/o creating a DB to keep init time short
		dbComp.CreateDataBase(dbConfig.DBCompositeTypes, dbConfig.DBModel)
	}
	// 3. Return (Reference) to Resource Allocation
	return dbComp
}

func (c *DBComponent) ConnectDataBase() {
	c.db = pg.Connect(&pg.Options{
		Addr:     c.config.Addr,
		User:     c.config.User,
		Password: c.config.Password,
		Database: c.config.Database,
	})
}

// CreateDataBase creates a DB for the given DB config
func (c *DBComponent) CreateDataBase(compositeTypes []interface{}, schema []interface{}) *pg.DB {
	mtd := "dbClient/CreateDataBase: "

	if compositeTypes != nil {
		compErr := c.CreateCompositeTypes(compositeTypes)
		if compErr != nil {
			log.Println(mtd + "Can't create or update DB composite types")
			log.Fatal(compErr)
		}
	}

	if schema != nil {
		dbCreateErr := c.CreateSchema(schema)
		if dbCreateErr != nil {
			log.Println(mtd + "Can't create or update DB schema")
			log.Fatal(dbCreateErr)
		}
	}
	return c.db
}

// CreateCompositeTypes creates composite types for the supplied array of structs.
// Must be called before creating the db schema to ensure all required types are present
func (c *DBComponent) CreateCompositeTypes(models []interface{}) error {

	db := c.db
	// Teardown happens in reverse order of creation.
	for i := len(models) - 1; i >= 0; i-- {
		err := db.Model(models[i]).DropComposite(&orm.DropCompositeOptions{
			IfExists: true,
			Cascade:  true,
		})
		if err != nil {
			return err
		}
	}

	// Type creation in actual order
	for _, model := range models {
		err := db.Model(model).CreateComposite(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateSchema creates database schema for the supplied array of structs.
func (c *DBComponent) CreateSchema(models []interface{}) error {

	db := c.db
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        !c.prod, // If prod, make persistent tables, else make temp tables
			IfNotExists: c.prod,  // if prod, only create tables if not exists already
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateManyToManyTables creates M2M relation tables for the supplied array of structs.
// Must be called after CreateSchema to ensure the presence of all tables
func (c *DBComponent) CreateManyToManyTables(models []interface{}) error {

	db := c.db
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Temp:        true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// PingDataBase returns true / false if the DB can be reached.
func (c *DBComponent) PingDataBase() (bool, error) {
	db := c.db
	err := db.Ping(db.Context())
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Shutdown closes the DB
func (c *DBComponent) Shutdown() error {
	err := c.db.Close()
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DB returns the DB
func (c *DBComponent) DB() *pg.DB {
	return c.db
}
