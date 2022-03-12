package database

const (
	ShipmentStatusCreatedStr              = "created"
	ShipmentStatusAllocatedStr            = "allocated"
	ShipmentStatusOnGoingToOriginStr      = "ongoing_to_origin"
	ShipmentStatusAtOriginStr             = "at_origin"
	ShipmentStatusOnGoingToDestinationStr = "ongoing_to_destination"
	ShipmentStatusAtDestinationStr        = "at_destination"
	ShipmentStatusCompletedStr            = "completed"
)

var (
	ShipmentStatusMap = map[string]int{
		ShipmentStatusCreatedStr:              0,
		ShipmentStatusAllocatedStr:            1,
		ShipmentStatusOnGoingToOriginStr:      2,
		ShipmentStatusAtOriginStr:             3,
		ShipmentStatusOnGoingToDestinationStr: 4,
		ShipmentStatusAtDestinationStr:        5,
		ShipmentStatusCompletedStr:            6,
	}

	ShipmentStatusIntMap = map[int]string{
		0: ShipmentStatusCreatedStr,
		1: ShipmentStatusAllocatedStr,
		2: ShipmentStatusOnGoingToOriginStr,
		3: ShipmentStatusAtOriginStr,
		4: ShipmentStatusOnGoingToDestinationStr,
		5: ShipmentStatusAtDestinationStr,
		6: ShipmentStatusCompletedStr,
	}
)

const (
	GetShipmentsDataQuery   = `SELECT * FROM shipments`
	InsertShipmentDataQuery = `INSERT INTO shipments (shipment_number, origin, destination, loading_date, status) VALUES (
		$1, $2, $3, $4, $5
	)`
	AllocateShipmentQuery = `UPDATE shipments SET 
		license_number = $1,
		driver_name = $2,
		update_time = now()
	WHERE shipment_number = $3`
	UpdateShipmentStatusQuery = `UPDATE shipments SET 
		status = $1,
		update_time = now()
	WHERE shipment_number = $2`

	GetShipmentsDataByShipmentsNumberQuery = GetShipmentsDataQuery + " where shipment_number in (arr) limit $1 order by create_time"
	GetShipmentsAllDataQuery               = GetShipmentsDataQuery + " order by create_time"
)
