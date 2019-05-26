package dao

import "database/sql"

type Models struct {
	ID            int64          `db:"id"`
	Cover         sql.NullString `db:"cover"`
	Name          sql.NullString `db:"name"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	NickNames     sql.NullString `db:"nicknames"`
	Birthday      sql.NullString `db:"birthday"`
	Constellation sql.NullString `db:"constellation"`
	Height        int64          `db:"height"`
	Dimensions    sql.NullString `db:"dimensions"`
	Cup           sql.NullString `db:"cup"`
	Address       sql.NullString `db:"address"`
	Jobs          sql.NullString `db:"jobs"`
	Interest      sql.NullString `db:"interest"`
	More          sql.NullString `db:"more"`
	Tags          sql.NullString `db:"tags"`
}
