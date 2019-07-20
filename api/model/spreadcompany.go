package dbModel

// `id` int(20) NOT NULL,
// `uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
// `companyName` varchar(200) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
// `companyCode` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
// `contact` varchar(100) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
// `legalPerson` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NULL DEFAULT NULL,
// `createdAt` datetime(0) NOT NULL,
// `updateAt` datetime(0) NOT NULL,

type SpreadCompany struct {
	Id          int64  `json:id db:"id"`
	Uuid        string `json:"uuid" db:"uuid"`
	CompanyName string `json:companyName db:"companyName"`
	CompanyCode string `json:companyCode db:"companyCode"`
	Contact     string `json:contact db:"contact"`
	LegalPerson string `json:legalPerson db:"legalPerson"`
	CreatedAt   string `json:createdAt db:"createdAt"`
	UpdateAt    string `json:updateAt db:"updateAt"`
}
