package idgen

import (
	"strconv"

	"github.com/sony/sonyflake"

	"hotpu.cn/xkefu/common/config"
)

var (
	generator *IDGenerator
)

func Get() (*IDGenerator, error) {
	if generator != nil {
		return generator, nil
	}
	return nil, config.ErrIDGeneratorInitialized
}

type IDGenerator struct {
	*sonyflake.Sonyflake
}

func (g *IDGenerator) NextID() (uint64, error) {
	return g.Sonyflake.NextID()
}

func (g *IDGenerator) NextIDString() (string, error) {
	v, err := g.Sonyflake.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(v, 10), nil
}

func InitFromConfig() error {
	st := sonyflake.Settings{}

	// use the connector's worker id as the machine id
	sf := sonyflake.NewSonyflake(st)
	gen := &IDGenerator{sf}

	generator = gen
	return nil
}
