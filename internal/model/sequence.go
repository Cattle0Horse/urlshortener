package model

// Sequence 序列号表
// 用于数据库号段发号器
type Sequence struct {
	Model
	Name     string `json:"name" gorm:"type:varchar(500);not null;uniqueIndex:idx_sequence_name;comment:序列名称"`
	Sequence int64  `json:"sequence" gorm:"type:bigint;not null;comment:当前序列值"`
	Version  int64  `json:"version" gorm:"type:bigint;comment:乐观锁版本号"`
}
