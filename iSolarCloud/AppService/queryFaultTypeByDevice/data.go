package queryFaultTypeByDevice

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/faultService/queryFaultTypeByDevice"
const Disabled = false
const EndPointName = "AppService.queryFaultTypeByDevice"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent          GoStruct.GoStructParent   `json:"-" DataTable:"true" DataTableSortOn:"FaultTypeCode"`

	FaultTypeCode     valueTypes.Integer   `json:"fault_type_code"`
	FaultTypeCodeList []valueTypes.Integer `json:"fault_type_code_list"`
	FaultTypeCode2    valueTypes.Integer   `json:"faulttypecode" PointId:"fault_type_code2"`
	DevFaultTypeCode  valueTypes.String    `json:"dev_fault_type_code"`
	FaultValue        valueTypes.String    `json:"faultvalue" PointId:"fault_value"`
	IsAllowOwnerView  valueTypes.Bool      `json:"is_allow_owner_view"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
