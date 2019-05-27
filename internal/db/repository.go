package db

import (
	"cadence-service/internal/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLRepository -
type MySQLRepository struct {
	db *sql.DB
}

// NewMySQLRepository - constructor
func NewMySQLRepository(url string) (*MySQLRepository, error) {

	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	// if ping doesnt work return nil
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &MySQLRepository{
		db,
	}, nil

}

// Handler - the implementation
type Handler interface {
	ReadOne(name string) (*model.DatastoreModel, error)
	ReadAll() ([]*model.DatastoreModel, error)
	Insert(player *model.DatastoreModel) error
	Delete(ID string) error
}

// ReadOne - returns a player from the db for the ID
func (r *MySQLRepository) ReadOne(name string) (*model.DatastoreModel, error) {

	fmt.Printf("Reading %s ", name)

	resultModel := model.DatastoreModel{}
	rows,err := r.db.Query("select name,team,number from player_info where name = $1",name)

	if err != nil{

		fmt.Errorf("cannot get rows  - error %v" , err)
		return nil,err
	}
	// close the connection at the end
	defer rows.Close()

	for rows.Next(){


		 rows.Scan(resultModel.Name,resultModel.Number,resultModel.Team)

		 if err != nil {
		 	fmt.Errorf("cannot marshall datastore model %v" , err)

		 }

		 fmt.Printf("Name %s , Number %d , Team %s" , resultModel.Name , resultModel.Number , resultModel.Team)
	}

	// check for the rows
	err = rows.Err()
	if err != nil {
		return nil,err
	}

	return &resultModel, nil
}

// ReadAll - returns all the players
func (r *MySQLRepository) ReadAll() ([]*model.DatastoreModel, error) {

	return nil, nil
}

// Insert - inserts a player
func (r *MySQLRepository) Insert(player *model.DatastoreModel) error {

	return nil
}

// Delete - deletes a player
func (r *MySQLRepository) Delete(ID string) error {

	return nil
}
