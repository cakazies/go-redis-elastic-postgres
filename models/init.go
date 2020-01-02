package models

type (
	// Provinces struct for table province
	Province struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	// Regency struct for table regency
	Regency struct {
		ID         string `json:"id"`
		ProvinceID string `json:"province_id"`
		Name       string `json:"name"`
	}

	// District struct for table district
	District struct {
		ID        string `json:"id"`
		RegencyID string `json:"regency_id"`
		Name      string `json:"name"`
	}

	// Village struct for table village
	Village struct {
		ID         string `json:"id"`
		DistrictID string `json:"district_id"`
		Name       string `json:"name"`
	}
)
