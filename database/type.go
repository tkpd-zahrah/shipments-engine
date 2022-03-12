package database

import "time"

type (
	Shipment struct {
		ShipmentNumber string    `json:"shipment_number"`
		LicenseNumber  string    `json:"license_number"`
		DriverName     string    `json:"driver_name"`
		Origin         string    `json:"origin"`
		Destination    string    `json:"destination"`
		LoadingDate    time.Time `json:"loading_date"`
		Status         int       `json:"status"`
	}

	AddShipmentParam struct {
		ShipmentNumber string `json:"shipment_number"`
		Origin         string `json:"origin"`
		Destination    string `json:"destination"`
		LoadingDate    string `json:"loading_date"`
	}

	AllocationParam struct {
		ShipmentNumber     string `json:"shipment_number"`
		TruckLicenseNumber string `json:"truck_license_number"`
		DriverName         string `json:"driver_name"`
	}
)
