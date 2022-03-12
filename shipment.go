package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zahrahfebriani/shipments-engine/database"
)

func GetShipmentsData(req GetShipmentsDataRequest) (GetShipmentsDataResponse, error) {
	shipmentsData, err := rsc.GetShipmentsData(req.ShipmentNumbers)
	if err != nil {
		log.Println("Failed GetShipmentsData", err.Error())
		return GetShipmentsDataResponse{
			Result: ResultStatus{
				Status:  "ERROR",
				Code:    "100001",
				Message: err.Error(),
			},
		}, err
	}

	return generateGetShipmentsDataResponse(shipmentsData), nil
}

func generateGetShipmentsDataResponse(shipmentsData []database.Shipment) GetShipmentsDataResponse {
	shipments := make([]Shipment, len(shipmentsData))
	for _, data := range shipmentsData {
		shipments = append(shipments, Shipment{
			ShipmentNumber: data.ShipmentNumber,
			LicenseNumber:  data.LicenseNumber,
			DriverName:     data.DriverName,
			Origin:         data.Origin,
			Destination:    data.Destination,
			LoadingDate:    data.LoadingDate,
			Status:         database.ShipmentStatusIntMap[data.Status],
		})
	}
	return GetShipmentsDataResponse{
		Result: ResultStatus{
			Status: "OK",
			Code:   "200",
		},
		Data: DataShipment{
			Shipments: shipments,
		},
	}
}

func AddShipmentData(req AddShipmentRequest) (AddShipmentResponse, error) {
	shipmentNumber := "Kargo" + fmt.Sprint(time.Now().Unix())

	err := rsc.InsertShipmentData(database.AddShipmentParam{
		ShipmentNumber: shipmentNumber,
		Origin:         req.Origin,
		Destination:    req.Destination,
		LoadingDate:    req.LoadingDate.Format("YYYY-MM-DD"),
	})
	if err != nil {
		log.Println("Failed AddShipmentData", err.Error())
		return AddShipmentResponse{
			Result: ResultStatus{
				Status:  "ERROR",
				Code:    "100002",
				Message: err.Error(),
			},
		}, err
	}

	return AddShipmentResponse{
		Result: ResultStatus{
			Status: "OK",
			Code:   "200",
		},
		ShipmentNumber: shipmentNumber,
	}, nil
}
