package node

type NodeInfo struct {
	Ips						[]HostInfo	`json:"ips"`
	IdcId                   int64		`json:"idc_id"`
	CityId                  int64		`json:"city_id"`
	AreaId                  int64		`json:"area_id"`
}