package queryDeviceList

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/devService/queryDeviceList"
const Disabled = false
const EndPointName = "AppService.queryDeviceList"

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []Device `json:"pageList" PointId:"devices"`

	DevCountByStatusMap struct {
		FaultCount   valueTypes.Count `json:"fault_count" PointId:"fault_count" PointUpdateFreq:"UpdateFreqTotal"`
		OfflineCount valueTypes.Count `json:"offline_count" PointId:"offline_count" PointUpdateFreq:"UpdateFreqTotal"`
		RunCount     valueTypes.Count `json:"run_count" PointId:"run_count" PointUpdateFreq:"UpdateFreqTotal"`
		WarningCount valueTypes.Count `json:"warning_count" PointId:"warning_count" PointUpdateFreq:"UpdateFreqTotal"`
	} `json:"dev_count_by_status_map" PointId:"device_status_count"`
	DevCountByTypeMap map[string]valueTypes.Integer `json:"dev_count_by_type_map" PointId:"device_type_count" PointUpdateFreq:"UpdateFreqBoot"`
	DevTypeDefinition map[string]valueTypes.String  `json:"dev_type_definition" PointId:"device_types" PointUpdateFreq:"UpdateFreqBoot"` // DataTable:"true"`

	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

// DevCountByTypeMap struct {
// 	One4 valueTypes.Integer `json:"14"`
// 	Two2 valueTypes.Integer `json:"22"`
// } `json:"dev_count_by_type_map"`
// DevTypeDefinition struct {
// 	One    valueTypes.String `json:"1"`
// 	One0   valueTypes.String `json:"10"`
// 	One1   valueTypes.String `json:"11"`
// 	One2   valueTypes.String `json:"12"`
// 	One3   valueTypes.String `json:"13"`
// 	One4   valueTypes.String `json:"14"`
// 	One5   valueTypes.String `json:"15"`
// 	One6   valueTypes.String `json:"16"`
// 	One7   valueTypes.String `json:"17"`
// 	One8   valueTypes.String `json:"18"`
// 	One9   valueTypes.String `json:"19"`
// 	Two0   valueTypes.String `json:"20"`
// 	Two1   valueTypes.String `json:"21"`
// 	Two2   valueTypes.String `json:"22"`
// 	Two3   valueTypes.String `json:"23"`
// 	Two4   valueTypes.String `json:"24"`
// 	Two5   valueTypes.String `json:"25"`
// 	Two6   valueTypes.String `json:"26"`
// 	Two8   valueTypes.String `json:"28"`
// 	Two9   valueTypes.String `json:"29"`
// 	Three  valueTypes.String `json:"3"`
// 	Three0 valueTypes.String `json:"30"`
// 	Three1 valueTypes.String `json:"31"`
// 	Three2 valueTypes.String `json:"32"`
// 	Three3 valueTypes.String `json:"33"`
// 	Three4 valueTypes.String `json:"34"`
// 	Three5 valueTypes.String `json:"35"`
// 	Three6 valueTypes.String `json:"36"`
// 	Three7 valueTypes.String `json:"37"`
// 	Three8 valueTypes.String `json:"38"`
// 	Three9 valueTypes.String `json:"39"`
// 	Four   valueTypes.String `json:"4"`
// 	Four0  valueTypes.String `json:"40"`
// 	Four1  valueTypes.String `json:"41"`
// 	Four2  valueTypes.String `json:"42"`
// 	Four3  valueTypes.String `json:"43"`
// 	Four4  valueTypes.String `json:"44"`
// 	Four5  valueTypes.String `json:"45"`
// 	Four6  valueTypes.String `json:"46"`
// 	Four7  valueTypes.String `json:"47"`
// 	Four8  valueTypes.String `json:"48"`
// 	Five   valueTypes.String `json:"5"`
// 	Five0  valueTypes.String `json:"50"`
// 	Six    valueTypes.String `json:"6"`
// 	Seven  valueTypes.String `json:"7"`
// 	Eight  valueTypes.String `json:"8"`
// 	Nine   valueTypes.String `json:"9"`
// 	Nine9  valueTypes.String `json:"99"`
// } `json:"dev_type_definition"`

type Device struct {
	GoStruct GoStruct.GoStruct `json:"-" PointIdFrom:"PsKey" PointIdReplace:"true" PointDeviceFrom:"PsKey"`

	PsKey      valueTypes.PsKey   `json:"ps_key" PointId:"ps_key" PointUpdateFreq:"UpdateFreqBoot"`
	PsId       valueTypes.PsId    `json:"ps_id" PointId:"ps_id" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceType valueTypes.Integer `json:"device_type" PointId:"device_type" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceCode valueTypes.Integer `json:"device_code" PointId:"device_code" PointUpdateFreq:"UpdateFreqBoot"`
	ChannelId  valueTypes.Integer `json:"chnnl_id" PointId:"channel_id" PointUpdateFreq:"UpdateFreqBoot"`
	Sn         valueTypes.String  `json:"sn" PointId:"sn" PointName:"Serial Number" PointUpdateFreq:"UpdateFreqBoot"`

	AlarmCount              valueTypes.Count   `json:"alarm_count" PointId:"alarm_count" PointUpdateFreq:"UpdateFreqTotal"`
	CommandStatus           valueTypes.Integer `json:"command_status" PointId:"command_status" PointUpdateFreq:"UpdateFreqInstant"`
	ComponentAmount         valueTypes.Integer `json:"component_amount" PointId:"component_amount"`
	DataFlag                valueTypes.Integer `json:"data_flag" PointId:"data_flag" PointUpdateFreq:"UpdateFreqBoot"`
	DataFlagDetail          valueTypes.Integer `json:"data_flag_detail" PointId:"data_flag_detail"`
	DeviceArea              valueTypes.Integer `json:"device_area" PointId:"device_area" PointUpdateFreq:"UpdateFreqBoot"` // References UUID and referenced by UUIDIndexCode
	DeviceAreaName          valueTypes.String  `json:"device_area_name" PointId:"device_area_name" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceId                valueTypes.Integer `json:"device_id" PointId:"device_id" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceModelCode         valueTypes.String  `json:"device_model_code" PointId:"device_model_code" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceModelId           valueTypes.Integer `json:"device_model_id" PointId:"device_model_id" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceName              valueTypes.String  `json:"device_name" PointId:"device_name" PointUpdateFreq:"UpdateFreqBoot"`
	DeviceStatus            valueTypes.Bool    `json:"device_status" PointId:"device_status" PointUpdateFreq:"UpdateFreqInstant"`
	FaultCount              valueTypes.Count   `json:"fault_count" PointId:"fault_count" PointUpdateFreq:"UpdateFreqTotal"`
	FaultStatus             valueTypes.String  `json:"fault_status" PointId:"fault_status" PointUpdateFreq:"UpdateFreqInstant"`
	FunctionEnum            valueTypes.String  `json:"function_enum" PointId:"function_enum" PointUpdateFreq:"UpdateFreqInstant"`
	InstallerAlarmCount     valueTypes.Count   `json:"installer_alarm_count" PointId:"installer_alarm_count" PointUpdateFreq:"UpdateFreqTotal"`
	InstallerDevFaultStatus valueTypes.Integer `json:"installer_dev_fault_status" PointId:"installer_dev_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
	InstallerFaultCount     valueTypes.Count   `json:"installer_fault_count" PointId:"installer_fault_count" PointUpdateFreq:"UpdateFreqTotal"`
	InverterModelType       valueTypes.Integer `json:"inverter_model_type" PointId:"inverter_model_type" PointUpdateFreq:"UpdateFreqBoot"`
	IsDeveloper             valueTypes.Bool    `json:"is_developer" PointId:"is_developer" PointUpdateFreq:"UpdateFreqBoot"`
	IsG2point5Module        valueTypes.Bool    `json:"is_g2point5_module" PointId:"is_g2point5_module" PointUpdateFreq:"UpdateFreqBoot"`
	IsInit                  valueTypes.Bool    `json:"is_init" PointId:"is_init" PointUpdateFreq:"UpdateFreqBoot"`
	IsSecond                valueTypes.Bool    `json:"is_second" PointId:"is_second" PointUpdateFreq:"UpdateFreqBoot"`
	IsSupportParamset       valueTypes.Bool    `json:"is_support_paramset" PointId:"is_support_paramset" PointUpdateFreq:"UpdateFreqBoot"`
	NodeTimestamps          interface{}        `json:"node_timestamps" PointId:"node_timestamps"`
	OwnerAlarmCount         valueTypes.Count   `json:"owner_alarm_count" PointId:"owner_alarm_count" PointUpdateFreq:"UpdateFreqTotal"`
	OwnerDevFaultStatus     valueTypes.Integer `json:"owner_dev_fault_status" PointId:"owner_dev_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
	OwnerFaultCount         valueTypes.Count   `json:"owner_fault_count" PointId:"owner_fault_count" PointUpdateFreq:"UpdateFreqTotal"`
	Points                  interface{}        `json:"points" PointId:"points"`
	RelState                valueTypes.Integer `json:"rel_state" PointId:"rel_state" PointUpdateFreq:"UpdateFreqInstant"`
	StringAmount            valueTypes.Integer `json:"string_amount" PointId:"string_amount"`
	TypeName                valueTypes.String  `json:"type_name" PointId:"type_name" PointUpdateFreq:"UpdateFreqBoot"`
	UnitName                valueTypes.String  `json:"unit_name" PointId:"unit_name" PointUpdateFreq:"UpdateFreqBoot"`
	UUID                    valueTypes.Integer `json:"uuid" PointId:"uuid" PointUpdateFreq:"UpdateFreqBoot"`                       // Referenced by DeviceArea
	UUIDIndexCode           valueTypes.String  `json:"uuid_index_code" PointId:"uuid_index_code" PointUpdateFreq:"UpdateFreqBoot"` // Referenced by DeviceArea

	PointData      []PointStruct `json:"point_data" PointId:"data" PointIdReplace:"true" DataTable:"true"` // PointIdFromChild:"PointId" PointIdReplace:"true" PointId:"data" DataTable:"true"`
	PsTimezoneInfo struct {
		IsDst    valueTypes.Bool   `json:"is_dst" PointUpdateFreq:"UpdateFreqInstant"`
		TimeZone valueTypes.String `json:"time_zone" PointUpdateFreq:"UpdateFreqInstant"`
	} `json:"psTimezoneInfo" PointId:"ps_timezone_info"`
}

type PointStruct struct {
	GoStruct GoStruct.GoStruct `json:"-" PointIdFrom:"PointId" PointIdReplace:"true" PointTimestampFrom:"TimeStamp" PointDeviceFromParent:"PsKey"`
	// GoStruct               GoStruct.GoStruct   `json:"-" PointDeviceFromParent:"PsKey"`

	TimeStamp        valueTypes.DateTime `json:"time_stamp" PointUpdateFreq:"UpdateFreq5Mins" PointNameDateFormat:"DateTimeLayout"`
	PointId          valueTypes.PointId  `json:"point_id" PointUpdateFreq:"UpdateFreqBoot"`
	PointName        valueTypes.String   `json:"point_name" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	Value            valueTypes.Float    `json:"value" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUnitFrom:"Unit" PointVariableUnit:"true" PointUpdateFreq:"UpdateFreq5Mins"`
	PointSign        valueTypes.String   `json:"point_sign" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	Unit             valueTypes.String   `json:"unit" PointUpdateFreq:"UpdateFreqBoot"`
	ValueDescription valueTypes.String   `json:"value_description" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	OrderId          valueTypes.Integer  `json:"order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupId     valueTypes.Integer  `json:"point_group_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupName   valueTypes.String   `json:"point_group_name" PointUpdateFreq:"UpdateFreqBoot"`
	Relate           valueTypes.Integer  `json:"relate" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	CodeId           valueTypes.Integer  `json:"code_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`

	CodeIdOrderId          valueTypes.String   `json:"code_id_order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	CodeName               valueTypes.String   `json:"code_name" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	DevPointLastUpdateTime valueTypes.DateTime `json:"dev_point_last_update_time" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreq5Mins" PointNameDateFormat:"DateTimeLayout"`
	IsPlatformDefaultUnit  valueTypes.Bool     `json:"is_platform_default_unit" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	IsShow                 valueTypes.Bool     `json:"is_show" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	OrderNum               valueTypes.Integer  `json:"order_num" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupIdOrderId    valueTypes.Integer  `json:"point_group_id_order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	ValIsFixed             valueTypes.Bool     `json:"val_is_fixd" PointId:"value_is_fixed" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	ValidSize              valueTypes.Integer  `json:"valid_size" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *ResultData) GetDataByName(name string) []PointStruct {
	var ret []PointStruct
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.DeviceName.Value() != name {
				continue
			}
			ret = p.PointData
			break
		}
	}
	return ret
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))

		// // Used for virtual entries.
		// // 0 - sungrow_battery_charging_power
		// var PVPowerToBattery
		// // sensor.sungrow_battery_discharging_power
		// var BatteryPowerToLoad
		// // 0 - sensor.sungrow_total_export_active_power
		// var PVPowerToGrid
		// // sensor.sungrow_purchased_power
		// var GridPowerToLoad
		// // 0 - sensor.sungrow_daily_battery_charging_energy_from_pv
		// var YieldBatteryCharge
		// // var DailyBatteryChargingEnergy
		// // sensor.sungrow_daily_battery_discharging_energy
		// var DailyBatteryDischargingEnergy
		// // 0 - sensor.sungrow_daily_feed_in_energy_pv
		// var YieldFeedIn
		// // sensor.sungrow_daily_purchased_energy
		// var DailyPurchasedEnergy
		// var PVPower
		// var LoadPower
		// var YieldSelfConsumption
		// // var DailyFeedInEnergy
		// var TotalPvYield
		// var DailyTotalLoad
		// var TotalEnergyConsumption
		for _, device := range e.Response.ResultData.PageList {
			epp := GoStruct.NewEndPointPath("virtual", device.PsKey.String())
			deviceId := device.PsKey.String()
			if device.PsKey.String() == "" {
				epp = GoStruct.NewEndPointPath("virtual", device.PsId.String())
				deviceId = device.PsId.String()
			}
			// Points are embedded within []PointStruct. So manually add virtuals instead of using the structure.

			for _, point := range device.PointData {
				name := point.PointId.String()
				foo := entries.CopyPointFromName(name, epp, name, point.PointName.String())
				foo.Value.Reset()
				foo.Value.AddFloat("", point.Unit.String(), "", point.Value.Value())
				// foo.SetUnit(point.Unit.String())
				foo.Value.SetDeviceId(deviceId)
				foo.DataStructure.PointGroupName = point.PointGroupName.String()
				foo.DataStructure.PointDevice = deviceId
				foo.DataStructure.ValueType = foo.Value.TypeValue
				foo.DataStructure.PointUnit = point.Unit.String()
				foo.DataStructure.PointTimestamp = point.TimeStamp.Time
				foo.IsOk = true
				// fmt.Printf("%s.%s -> %s\n", epp, name, foo.DataStructure.PointDevice)
			}
		}

		e.GetEnergyStorageSystem(entries)
		// e.GetCommunicationModule(entries)
		// e.GetBattery(entries)
	}

	return entries
}

func (e *EndPoint) GetEnergyStorageSystem(entries api.DataMap) {
	for range Only.Once {
		/*
			PVPower				- TotalDcPower
			PVPowerToBattery	- BatteryChargingPower
			PVPowerToLoad		- TotalDcPower - BatteryChargingPower - TotalExportActivePower
			PVPowerToGrid		- TotalExportActivePower

			LoadPower			- TotalLoadActivePower
			BatteryPowerToLoad	- BatteryDischargingPower
			BatteryPowerToGrid	- ?

			GridPower			- lowerUpper(PVPowerToGrid, GridPowerToLoad)
			GridPowerToLoad		- PurchasedPower
			GridPowerToBattery	- ?

			YieldSelfConsumption	- DailyLoadEnergyConsumptionFromPv
			YieldBatteryCharge		- DailyBatteryChargingEnergyFromPv
			YieldFeedIn				- DailyFeedInEnergyPv
		*/

		for _, device := range e.Response.ResultData.PageList {
			if !device.DeviceType.Match(api.DeviceNameEnergyStorageSystem) {
				// Only looking for a Solar Storage System.
				continue
			}
			epp := GoStruct.NewEndPointPath("virtual", device.PsKey.String())
			if device.PsKey.String() == "" {
				epp = GoStruct.NewEndPointPath("virtual", device.PsId.String())
			}
			// Points are embedded within []PointStruct. So manually add virtuals instead of using the structure.
			e.SetBatteryPoints(epp, entries)
			e.SetPvPoints(epp, entries)
			e.SetGridPoints(epp, entries)
			e.SetLoadPoints(epp, entries)
		}
	}
}

func (e *EndPoint) SetBatteryPoints(epp GoStruct.EndPointPath, entries api.DataMap) {
	for range Only.Once {
		// /////////////////////////////////////////////////////// //
		// Battery Power
		batteryChargePower := entries.CopyPointFromName(epp.AddString("p13126"), epp, "battery_charge_power", "Battery Charge Power (p13126)")
		// batteryChargePower.DataStructure.PointIcon = "mdi:battery"
		batteryChargePower.DataStructure.PointUpdateFreq = GoStruct.UpdateFreq5Mins

		batteryDischargePower := entries.CopyPointFromName(epp.AddString("p13150"), epp, "battery_discharge_power", "Battery Discharge Power (p13150)")
		// batteryDischargePower.DataStructure.PointIcon = "mdi:battery"
		batteryDischargePower.DataStructure.PointUpdateFreq = GoStruct.UpdateFreq5Mins

		batteryPower := entries.CopyPoint(batteryChargePower, epp, "battery_power", "Battery Power (Calc)")
		batteryPower.SetValue(entries.LowerUpper(batteryDischargePower, batteryChargePower))
		// batteryPower.DataStructure.PointIcon = "mdi:battery"
		batteryPower.DataStructure.PointUpdateFreq = GoStruct.UpdateFreq5Mins

		batteryPowerActive := entries.CopyPoint(batteryPower, epp, "battery_power_active", "Battery Power Active (Calc)")
		// batteryPowerActive.DataStructure.PointIcon = "mdi:battery"
		_ = entries.MakeState(batteryPowerActive)
		batteryPowerActive.DataStructure.PointUpdateFreq = GoStruct.UpdateFreq5Mins

		// /////////////////////////////////////////////////////// //
		batteryDischargeEnergy := entries.CopyPointFromName(epp.AddString("p13029"), epp, "battery_discharge_energy", "Battery Discharge Energy (p13029)")
		// batteryDischargeEnergy.DataStructure.PointIcon = "mdi:battery"
		batteryDischargeEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		batteryChargeEnergy := entries.CopyPointFromName(epp.AddString("p13174"), epp, "battery_charge_energy", "Battery Charge Energy (p13174)")
		// batteryChargeEnergy.DataStructure.PointIcon = "mdi:battery"
		batteryChargeEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		batteryEnergy := entries.CopyPoint(batteryChargeEnergy, epp, "battery_energy", "Battery Energy (Calc)")
		batteryEnergy.SetValue(entries.LowerUpper(batteryChargeEnergy, batteryDischargeEnergy))
		// batteryEnergy.DataStructure.PointIcon = "mdi:battery"
		batteryEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		batteryEnergyActive := entries.CopyPoint(batteryEnergy, epp, "battery_energy_active", "Battery Energy Active (Calc)")
		// batteryEnergyActive.DataStructure.PointIcon = "mdi:battery"
		batteryEnergyActive.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		_ = entries.MakeState(batteryEnergyActive)

		dailyPvEnergy := entries.GetReflect(epp.AddString("p13112"))
		batteryChargeEnergyPercent := entries.CopyPoint(dailyPvEnergy, epp, "battery_charge_energy_percent", "Battery Charge Percent (Calc)")
		batteryChargeEnergyPercent.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		batteryChargeEnergyPercent.SetValue(entries.GetPercent(batteryChargeEnergy, dailyPvEnergy, 1))
		batteryChargeEnergyPercent.SetUnit("%")
		// batteryChargeEnergyPercent.DataStructure.PointIcon = "mdi:battery"

		// /////////////////////////////////////////////////////// //
		batteryToLoadPower := entries.CopyPoint(batteryDischargePower, epp, "battery_to_load_power", "Battery To Load Power (Calc)")
		// batteryToLoadPower.DataStructure.PointIcon = "mdi:battery"
		batteryToLoadPower.DataStructure.PointUpdateFreq = GoStruct.UpdateFreq5Mins

		batteryToLoadPowerActive := entries.CopyPoint(batteryToLoadPower, epp, "battery_to_load_power_active", "Battery To Load Power Active (Calc)")
		// batteryToLoadPowerActive.DataStructure.PointIcon = "mdi:battery"
		batteryToLoadPowerActive.DataStructure.PointUpdateFreq = GoStruct.UpdateFreq5Mins
		_ = entries.MakeState(batteryToLoadPowerActive)

		batteryToGridPowerActive := entries.CopyPoint(batteryDischargePower, epp, "battery_to_grid_power_active", "Battery To Grid Power Active (Calc)")
		batteryToGridPowerActive.SetValue(0.0)
		// batteryToGridPowerActive.DataStructure.PointIcon = "mdi:battery"
		batteryToGridPowerActive.DataStructure.PointUpdateFreq = GoStruct.UpdateFreq5Mins
		_ = entries.MakeState(batteryToGridPowerActive)
	}
}

func (e *EndPoint) SetPvPoints(epp GoStruct.EndPointPath, entries api.DataMap) {
	for range Only.Once {
		// /////////////////////////////////////////////////////// //
		// PV Power
		pvPower := entries.CopyPointFromName(epp.AddString("p13003"), epp, "pv_power", "Pv Power (p13003)")

		pvPowerActive := entries.CopyPoint(pvPower, epp, "pv_power_active", "Pv Power Active (p13003)")
		_ = entries.MakeState(pvPowerActive)

		pvToGridPower := entries.CopyPointFromName(epp.AddString("p13121"), epp, "pv_to_grid_power", "Pv To Grid Power (p13121)")

		PvToGridPowerActive := entries.CopyPoint(pvToGridPower, epp, "pv_to_grid_power_active", "Pv To Grid Power Active (p13121)")
		_ = entries.MakeState(PvToGridPowerActive)

		pvToBatteryPower := entries.CopyPointFromName(epp.AddString("p13126"), epp, "pv_to_battery_power", "Pv To Battery Power (p13126)")

		pvToBatteryPowerActive := entries.CopyPoint(pvToBatteryPower, epp, "pv_to_battery_power_active", "Pv To Battery Power Active (p13126)")
		pvToBatteryPowerActive.SetValue(pvToBatteryPower.GetValueFloat())
		_ = entries.MakeState(pvToBatteryPowerActive)

		pvToLoadPower := entries.CopyPoint(pvPower, epp, "pv_to_load_power", "Pv To Load Power (Calc)")
		one := pvPower.GetValueFloat()
		two := pvToBatteryPower.GetValueFloat()
		three := pvToGridPower.GetValueFloat()
		pvToLoadPower.SetValue(one - two - three)
		pvToLoadPower.SetValuePrecision(3)

		pvToLoadPowerActive := entries.CopyPoint(pvToLoadPower, epp, "pv_to_load_power_active", "Pv To Load Power Active (Calc)")
		_ = entries.MakeState(pvToLoadPowerActive)

		pvDailyEnergy := entries.CopyPointFromName(epp.AddString("p13112"), epp, "pv_daily_energy", "Pv Daily Energy (p13112)")
		pvDailyEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		pvToGridEnergy := entries.CopyPointFromName(epp.AddString("p13173"), epp, "pv_to_grid_energy", "Pv To Grid Energy (p13173)")
		pvToGridEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		pvToGridEnergyPercent := entries.CopyPoint(pvDailyEnergy, epp, "pv_to_grid_energy_percent", "Pv To Grid Energy Percent (Calc)")
		pvToGridEnergyPercent.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		pvToGridEnergyPercent.SetValue(entries.GetPercent(pvToGridEnergy, pvDailyEnergy, 1))
		pvToGridEnergyPercent.SetUnit("%")

		pvToBatteryEnergy := entries.CopyPointFromName(epp.AddString("p13174"), epp, "pv_to_battery_energy", "Pv To Battery Energy (p13174)")
		pvToBatteryEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		// dailyPvEnergy(p13112) - pvToGridEnergy(p13173) - pvToBatteryEnergy(p13174)
		// WRONG!!! - p13112 (Pv Daily Energy) - p13122 (Daily Feed-in Energy) - p13174 (Daily Battery Charging Energy from PV)
		dailyFeedInEnergy := entries.GetReflect(epp.AddString("p13173"))
		batteryChargeEnergy := entries.GetReflect(epp.AddString("p13174"))
		selfConsumptionOfPv := entries.CopyPointFromName(epp.AddString("p13116"), epp, "pv_consumption_energy", "Pv Consumption Energy (Calc)")
		selfConsumptionOfPv.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		tmp1 := pvDailyEnergy.GetValueFloat() - dailyFeedInEnergy.GetValueFloat() - batteryChargeEnergy.GetValueFloat()
		selfConsumptionOfPv.SetValue(tmp1)
		selfConsumptionOfPv.SetValuePrecision(3)

		selfConsumptionOfPvPercent := entries.CopyPointFromName(epp.AddString("p13116"), epp, "pv_consumption_energy_percent", "Pv Consumption Energy Percent (Calc)")
		selfConsumptionOfPvPercent.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		selfConsumptionOfPvPercent.SetValue(entries.GetPercent(selfConsumptionOfPv, pvDailyEnergy, 1))
		selfConsumptionOfPvPercent.SetUnit("%")

		// WRONG!!! - pvToLoadPercent := entries.CopyPointFromName(epp.AddString("p13144"), epp, "pv_to_load_energy_percent", "Pv To Load Energy Percent (p13144)")
		// WRONG!!! - pvToLoadEnergy := entries.CopyPointFromName(epp.AddString("p13116"), epp, "pv_to_load_energy", "Pv To Load Energy (p13116)")
		gridToLoadEnergy := entries.GetReflect(epp.AddString("p13147"))
		pvToLoadEnergy := entries.CopyPointFromName(epp.AddString("p13116"), epp, "pv_to_load_energy", "Pv To Load Energy (p13116)")
		pvToLoadEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		tmp2 := pvToLoadEnergy.GetValueFloat() + gridToLoadEnergy.GetValueFloat()

		totalDailyEnergy := entries.CopyPoint(gridToLoadEnergy, epp, "total_daily_energy", "Total Daily Energy (Calc)")
		totalDailyEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		totalDailyEnergy.SetValue(tmp2)
		totalDailyEnergy.SetValuePrecision(3)

		pvToLoadEnergyPercent := entries.CopyPoint(pvDailyEnergy, epp, "pv_to_load_energy_percent", "Pv To Load Energy Percent (Calc)")
		pvToLoadEnergyPercent.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		pvToLoadEnergyPercent.SetValue(entries.GetPercent(pvToLoadEnergy, totalDailyEnergy, 1))
		pvToLoadEnergyPercent.SetUnit("%")

		gridToLoadDailyEnergy := entries.GetReflect(epp.AddString("p13147"))
		pvDailyEnergyPercent := entries.CopyPoint(totalDailyEnergy, epp, "pv_daily_energy_percent", "Pv Daily Energy Percent (Calc)")
		pvDailyEnergyPercent.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		dpe := totalDailyEnergy.GetValueFloat() - gridToLoadDailyEnergy.GetValueFloat()
		pvDailyEnergyPercent.SetValue(api.GetPercent(dpe, totalDailyEnergy.GetValueFloat(), 1))
		pvDailyEnergyPercent.SetUnit("%")

		pvEnergy := entries.CopyPointFromName(pvToLoadEnergy.PointId(), epp, "pv_energy", "Pv Energy (Calc)")
		// pvDailyYield := entries.GetReflect(pvSelfConsumption.PointId())
		pvEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		pvEnergy.SetValue(GoStruct.AddFloatValues(3, pvToLoadEnergy, pvToBatteryEnergy, pvToGridEnergy))

		// DailyFeedInEnergy - @TODO - This may differ from DailyFeedInEnergyPv
		// _ = entries.CopyPointFromName(epp.AddString("p13122"), epp, "pv_to_grid2", "")

		// TotalPvYield
		pcTotalEnergy := entries.CopyPointFromName(epp.AddString("p13134"), epp, "pv_total_energy", "Pv Total Energy (p13134)")
		pcTotalEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
	}
}

func (e *EndPoint) SetGridPoints(epp GoStruct.EndPointPath, entries api.DataMap) {
	for range Only.Once {
		gridToLoadPower := entries.CopyPointFromName(epp.AddString("p13149"), epp, "grid_to_load_power", "Grid To Load Power (p13149)")

		gridToLoadPowerActive := entries.CopyPoint(gridToLoadPower, epp, "grid_to_load_power_active", "Grid To Load Power Active (p13149)")
		_ = entries.MakeState(gridToLoadPowerActive)

		gridToLoadEnergy := entries.CopyPointFromName(epp.AddString("p13147"), epp, "grid_to_load_energy", "Grid To Load Energy (p13147)")
		gridToLoadEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		pvToLoadEnergy := entries.CopyPointFromName(epp.AddString("p13116"), epp, "pv_to_load_energy", "Pv To Load Energy (p13116)")
		pvToLoadEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		tmp2 := pvToLoadEnergy.GetValueFloat() + gridToLoadEnergy.GetValueFloat()

		totalLoadEnergy := entries.CopyPoint(gridToLoadEnergy, epp, "total_load_energy", "Total Load Energy (Calc)")
		totalLoadEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		totalLoadEnergy.SetValue(tmp2)
		totalLoadEnergy.SetValuePrecision(3)

		gridToLoadEnergyPercent := entries.CopyPoint(totalLoadEnergy, epp, "grid_to_load_energy_percent", "")
		gridToLoadEnergyPercent.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		gridToLoadEnergyPercent.SetValue(entries.GetPercent(gridToLoadEnergy, totalLoadEnergy, 1))
		gridToLoadEnergyPercent.SetUnit("%")

		pvToGridPower := entries.CopyPointFromName(epp.AddString("p13121"), epp, "pv_to_grid_power", "Pv To Grid Power (p13121)")
		gridPower := entries.CopyPoint(gridToLoadPower, epp, "grid_power", "Grid Power (Calc)")
		gridPower.SetValue(entries.LowerUpper(pvToGridPower, gridToLoadPower))

		gridPowerActive := entries.CopyPoint(gridPower, epp, "grid_power_active", "Grid Power Active (Calc)")
		_ = entries.MakeState(gridPowerActive)

		gridToBatteryPowerActive := entries.CopyPoint(gridToLoadPower, epp, "grid_to_battery_power_active", "Grid To Battery Power Active (Calc)")
		gridToBatteryPowerActive.SetValue(0.0)
		_ = entries.MakeState(gridToBatteryPowerActive)

		pvToGridEnergy := entries.CopyPointFromName(epp.AddString("p13173"), epp, "pv_to_grid_energy", "Pv To Grid Energy (p13173)")
		pvToGridEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		gridEnergy := entries.CopyPoint(pvToGridEnergy, epp, "grid_energy", "Grid Energy (Calc)")
		gridEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		gridEnergy.SetValue(entries.LowerUpper(pvToGridEnergy, gridToLoadEnergy))
	}
}

func (e *EndPoint) SetLoadPoints(epp GoStruct.EndPointPath, entries api.DataMap) {
	for range Only.Once {
		// Daily Load Energy Consumption

		gridToLoadEnergy := entries.CopyPointFromName(epp.AddString("p13147"), epp, "grid_to_load_energy", "Grid To Load Energy (p13147)")
		gridToLoadEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay

		pvToLoadEnergy := entries.CopyPointFromName(epp.AddString("p13116"), epp, "pv_to_load_energy", "Pv To Load Energy (p13116)")
		pvToLoadEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		tmp2 := pvToLoadEnergy.GetValueFloat() + gridToLoadEnergy.GetValueFloat()

		dailyTotalEnergy := entries.CopyPoint(gridToLoadEnergy, epp, "daily_total_energy", "Daily Total Energy (Calc)")
		dailyTotalEnergy.DataStructure.PointUpdateFreq = GoStruct.UpdateFreqDay
		dailyTotalEnergy.SetValue(tmp2)
		dailyTotalEnergy.SetValuePrecision(3)

		// Total Load Energy Consumption
		// _ = entries.CopyPointFromName(epp.AddString("p13130"), epp, "total_energy_consumption", "")

		loadPower := entries.CopyPointFromName(epp.AddString("p13119"), epp, "load_power", "Load Power (p13119)")

		loadPowerActive := entries.CopyPoint(loadPower, epp, "load_power_active", "Load Power Active (p13119)")
		_ = entries.MakeState(loadPowerActive)
	}
}
