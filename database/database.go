package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type ShipmentsDomain interface {
	GetShipmentsData(shipmentNumbers []string) ([]Shipment, error)
	InsertShipmentData(param AddShipmentParam) error
	AllocateShipment(param AllocationParam) error
	UpdateStatusShipment(status int, shipmentNumber string) error
}

type ShipmentResource struct {
	db   *sql.DB
	stmt *sql.Stmt
}

func replaceGetQuery(limit int, params []string, query string) string {
	strParam := ""
	for i, p := range params {
		strParam += p
		if i != len(params)-1 {
			strParam += ","
		}
	}

	res := strings.Replace(query, "arr", strParam, -1)
	res = strings.Replace(res, "$1", fmt.Sprint(limit), -1)
	return res
}

func InitShipment(database *sql.DB) *ShipmentResource {
	return &ShipmentResource{
		db: database,
	}
}

func (s *ShipmentResource) GetShipmentsData(shipmentNumbers []string, max int) ([]Shipment, error) {
	shipments := make([]Shipment, len(shipmentNumbers))

	if len(shipmentNumbers) > 0 && max > 0 {
		query := replaceGetQuery(max, shipmentNumbers, GetShipmentsDataByShipmentsNumberQuery)
		log.Println(query)
		rows, err := s.db.Query(query)
		if err != nil {
			return []Shipment{}, err
		}
		defer rows.Close()
		for rows.Next() {
			var shipment Shipment
			if err := rows.Scan(&shipment); err != nil {
				return []Shipment{}, err
			}
			shipments = append(shipments, shipment)
		}
	} else {
		rows, err := s.db.Query(GetShipmentsAllDataQuery)
		if err != nil {
			return []Shipment{}, err
		}
		defer rows.Close()
		for rows.Next() {
			var shipment Shipment
			if err := rows.Scan(&shipment); err != nil {
				return []Shipment{}, err
			}
			shipments = append(shipments, shipment)
		}
	}

	return shipments, nil
}

func (s *ShipmentResource) InsertShipmentData(param AddShipmentParam) error {
	if _, err := s.db.Exec(InsertShipmentDataQuery, param.ShipmentNumber, param.Origin, param.Destination, param.LoadingDate, ShipmentStatusMap[ShipmentStatusCreatedStr]); err != nil {
		return err
	}
	return nil
}

func (s *ShipmentResource) AllocateShipment(param AllocationParam) error {
	if _, err := s.db.Exec(AllocateShipmentQuery, param.TruckLicenseNumber, param.DriverName, param.ShipmentNumber); err != nil {
		return err
	}
	return nil
}

func (s *ShipmentResource) UpdateStatusShipment(status int, shipmentNumber string) error {
	if _, err := s.db.Exec(UpdateShipmentStatusQuery, status, shipmentNumber); err != nil {
		return err
	}
	return nil
}
