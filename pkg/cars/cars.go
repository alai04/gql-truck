package cars

import (
	"log"

	"github.com/alai04/gql-truck/graph/model"
	"github.com/alai04/gql-truck/pkg/db"
)

func ApprovedCars() []*model.TruckCar {
	stmt, err := db.Db.Prepare(
		`select C.id, C.plate_no, M.id, M.status_code, M.status_desc, M.spsj 
		from t_traffic_car C inner join t_traffic_main M on C.main_id = M.id 
		where M.status_code = 3`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cars []*model.TruckCar
	for rows.Next() {
		var car model.TruckCar
		var main model.TruckMain
		err := rows.Scan(&car.ID, &car.Plate, &main.ID, &main.StatusCode, &main.StatusDesc, &main.Spsj)
		if err != nil {
			log.Fatal(err)
		}
		car.Main = &main
		cars = append(cars, &car)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return cars
}

func Approved(plate string) bool {
	stmt, err := db.Db.Prepare(
		`select count(1) 
		from t_traffic_car C inner join t_traffic_main M on C.main_id = M.id 
		where M.status_code = 3 and C.plate_no = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(plate).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}
