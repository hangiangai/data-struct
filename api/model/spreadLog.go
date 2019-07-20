package dbModel

// `createAt` datetime(0) NOT NULL,
// `achievement` int(20) NULL DEFAULT NULL,
// `targetProcedure` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
// `targetContact` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
// `targetInfo` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
// `staffName` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
// `staffId` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
// `id` int(20) NOT NULL AUTO_INCREMENT,
// `uuid` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
// `update` datetime(0) NULL DEFAULT NULL,

type SpreadLog struct {
	Id              int64  `json:id db:"id"`
	Uuid            string `json:"uuid" db:"uuid"`
	Achievement     string `json:achievement db:"achievement"`
	TargetProcedure string `json:targetProcedure db:"targetProcedure"`
	TargetContact   string `json:targetContact db:"targetContact"`
	TargetInfo      string `json:targetInfo db:"targetInfo"`
	StaffName       string `json:staffName db:"staffName"`
	StaffId         string `json:staffId db:"staffId"`
	CreatedAt       string `json:"createdAt" db:"createdAt"`
	Update          string `json:update db:"update"`
}
