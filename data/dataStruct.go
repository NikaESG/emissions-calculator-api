package data

type ConfigStruct struct {
	// UserId   int64             `json:"user_id"`
	// UserName string            `json:"user_name"`
	Params map[string]string `json:"params"`
}

type ListConnectionsResp struct {
	Ret    int32
	Msg    string
	Config map[string]ConfigStruct `json:"config"`
}

type GetDataReq struct {
	TableName string `json:"table_name"`
	FieldName string `json:"field_name"`
}

type GetData struct {
	Ret  int32
	Msg  string
	Data []string `json:"data"`
}

// type Connector struct {
// 	CatalogId string
// }

// type GetUserConnectionsReq struct {
// 	UserId int64 `json:"user_id"`
// }

// type GetUserConnectionResp struct {
// 	Ret  int32
// 	Msg  string
// 	Data []string //Connector
// }

// type GetConnectorConfigReq struct {
// 	CatalogId string `json:"catalog_id"`
// }

// type GetConnectorConfigResp struct {
// 	Ret  int32
// 	Msg  string
// 	Data ConfigStruct
// }

type AddConnectionReq struct {
	// UserId    int64             `json:"user_id"`
	CatalogId string            `json:"catalog_id"`
	Data      map[string]string `json:"data"`
}

type AddConnectionResp struct {
	Ret int32
	Msg string
}

type UpdateConnectionReq struct {
	CatalogId string            `json:"catalog_id"`
	Data      map[string]string `json:"data"`
}

type UpdateConnectionResp struct {
	Ret int32
	Msg string
}

type DeleteConnectionReq struct {
	CatalogId string `json:"catalog_id"`
}

type DeleteConnectionResp struct {
	Ret int32
	Msg string
}

type TestConnectionReq struct {
	UserId    int64  `json:"user_id"`
	CatalogId string `json:"catalog_id"`
}

type TestConnectionResp struct {
	Ret int32
	Msg string
}
