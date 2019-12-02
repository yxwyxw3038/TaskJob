package model

type PublicResult struct {
	Code   string      `json:"code"`
	Data   interface{} `json:"data"`
	Reason string      `json:"reason"`
}

type CreatePdfListModel struct {
	Id          string `json:"Id"`
	MouldID     string `json:"MouldID"`
	StructureID string `json:"StructureID"`
	Type        int64  `json:"Type"`
}
type RequestInfoModel struct {
	Url          string           `json:"Url"`
	Action       string           `json:"Action"`
	IntervalTime int64            `json:"IntervalTime"`
	Parameter    []ParameterModel `json:"Parameter"`
}
type ParameterModel struct {
	Column string      `json:"Column"`
	Value  interface{} `json:"Value"`
}
