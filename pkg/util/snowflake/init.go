package snowflake

import "github.com/yitter/idgenerator-go/idgen"

func Init(workID uint16) {
	option := idgen.NewIdGeneratorOptions(workID)
	idgen.SetIdGenerator(option)
}
