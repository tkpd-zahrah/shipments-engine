package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zahrahfebriani/shipments-engine/database"
)

func GetShipmentsData(req GetShipmentsDataRequest) (GetShipmentsDataResponse, error) {
	shipmentsData, err := rsc.GetShipmentsData(req.ShipmentNumbers, req.Max)
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
	shipments := make([]Shipment, 0)
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
	log.Println("req", req)

	err := rsc.InsertShipmentData(database.AddShipmentParam{
		ShipmentNumber: shipmentNumber,
		Origin:         req.Origin,
		Destination:    req.Destination,
		LoadingDate:    req.LoadingDate.Format("2006-01-02"),
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

func AllocateShipment(req AllocationRequest) (AllocationResponse, error) {
	err := rsc.AllocateShipment(database.AllocationParam{
		ShipmentNumber:     req.ShipmentNumber,
		TruckLicenseNumber: req.TruckLicenseNumber,
		DriverName:         req.DriverName,
	})
	if err != nil {
		log.Println("Failed AllocateShipment", err.Error())
		return AllocationResponse{
			Result: ResultStatus{
				Status:  "ERROR",
				Code:    "100003",
				Message: err.Error(),
			},
		}, err
	}

	return AllocationResponse{
		Result: ResultStatus{
			Status: "OK",
			Code:   "200",
		},
	}, nil
}

func UpdateStatusShipment(req UpdateStatusShipmentRequest) (UpdateStatusShipmentResponse, error) {
	err := rsc.UpdateStatusShipment(database.ShipmentStatusMap[req.Status], req.ShipmentNumber)
	if err != nil {
		log.Println("Failed UpdateStatusShipment", err.Error())
		return UpdateStatusShipmentResponse{
			Result: ResultStatus{
				Status:  "ERROR",
				Code:    "100004",
				Message: err.Error(),
			},
		}, err
	}

	return UpdateStatusShipmentResponse{
		Result: ResultStatus{
			Status: "OK",
			Code:   "200",
		},
	}, nil
}
