package main

import "time"

type (
	ResultStatus struct {
		Status  string `json:"status"`
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	// GET SHIPMENTS DATA
	GetShipmentsDataRequest struct {
		ShipmentNumbers []string `json:"shipment_numbers"`
		Max             int      `json:"max"`
	}

	GetShipmentsDataResponse struct {
		Result ResultStatus `json:"result"`
		Data   DataShipment `json:"data"`
	}

	DataShipment struct {
		Shipments []Shipment `json:"shipments"`
	}

	Shipment struct {
		ShipmentNumber string    `json:"shipment_number"`
		LicenseNumber  string    `json:"license_number"`
		DriverName     string    `json:"driver_name"`
		Origin         string    `json:"origin"`
		Destination    string    `json:"destination"`
		LoadingDate    time.Time `json:"loading_date"`
		Status         string    `json:"status"`
	}

	// Add Shipment
	AddShipmentRequest struct {
		ShipmentNumber string    `json:"shipment_number"`
		Origin         string    `json:"origin"`
		Destination    string    `json:"destination"`
		LoadingDate    time.Time `json:"loading_date"`
	}

	AddShipmentResponse struct {
		Result         ResultStatus `json:"result"`
		ShipmentNumber string       `json:"shipment_number"`
	}

	AllocationParam struct {
		ShipmentNumber     string `json:"shipment_number"`
		TruckLicenseNumber string `json:"truck_license_number"`
		DriverName         string `json:"driver_name"`
	}
)
