package database

import (
	"database/sql"
	"errors"
)

type ShipmentsDomain interface {
	GetShipmentsData(shipmentNumbers []string) ([]Shipment, error)
	InsertShipmentData(param AddShipmentParam) error
	AllocateShipment(param AllocationParam) error
	UpdateStatusShipment(status int, shipmentNumber string) error
}

type ShipmentResource struct {
	db *sql.DB
}

// build query param
func buildQueryParam(param []interface{}) string {
	var res string
	for i, s := range param {
		res += s.(string)
		if i != len(param)-1 {
			res += ","
		}
	}
	return res
}

func InitShipment(database *sql.DB) *ShipmentResource {
	return &ShipmentResource{
		db: database,
	}
}

func (s *ShipmentResource) GetShipmentsData(shipmentNumbers []string) ([]Shipment, error) {
	shipments := make([]Shipment, len(shipmentNumbers))
	stmt, err := s.db.Prepare(GetShipmentsDataByShipmentsNumberQuery)
	if err != nil {
		return []Shipment{}, errors.New("failed to prepare statement")
	}

	queryParams := make([]interface{}, len(shipmentNumbers))
	for _, snumber := range shipmentNumbers {
		queryParams = append(queryParams, snumber)
	}

	rows, err := stmt.Query(buildQueryParam(queryParams))
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
	return shipments, nil
}

func (s *ShipmentResource) InsertShipmentData(param AddShipmentParam) error {
	if _, err := s.db.Exec(InsertShipmentDataQuery, param.ShipmentNumber, param.Origin, param.Destination, param.LoadingDate); err != nil {
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
