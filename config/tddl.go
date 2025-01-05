package config

type TDDL struct {
	Step     uint64 `yaml:"step" mapstructure:"step"`
	SeqName  string `yaml:"seq_name" mapstructure:"seq_name"`
	StartNum uint64 `yaml:"start_num" mapstructure:"start_num"`
}
