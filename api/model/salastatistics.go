package dbModel

import "time"

// `id` int(200) NOT NULL AUTO_INCREMENT,
// `uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
// `staffname` varchar(100) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
// `staffid` varchar(100) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
// `createdat` datetime(0) NOT NULL,
// `updateat` datetime(0) NOT NULL,

type SalaStatistics struct {
	Id        int64     `json:id db:"id"`
	Uuid      string    `json:uuid db:uuid`
	StaffName string    `json:staffname db:"staffname"`
	StaffId   string    `json:staffid db:"staffid"`
	CreatedAt time.Time `json:createdat db:"createdat`
	UpdateAt  time.Time `json:updateat db:"updateat"`
}
