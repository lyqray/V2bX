package conf

import (
	"fmt"
	"io"
	"os"

	"github.com/InazumaV/V2bX/common/json5"

	"encoding/json/v2"
)

const DefaultNodeRetryCount = 1
const DefaultNodeTimeout = 15

type Conf struct {
	LogConfig   LogConfig    `json:"Log"`
	CoresConfig []CoreConfig `json:"Cores"`
	NodeConfig  []NodeConfig `json:"Nodes"`
}

func New() *Conf {
	return &Conf{
		LogConfig: LogConfig{
			Level:  "info",
			Output: "",
		},
	}
}

func (p *Conf) LoadFromPath(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open config file error: %s", err)
	}
	defer f.Close()

	reader := json5.NewTrimNodeReader(f)
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("read config file error: %s", err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		return fmt.Errorf("unmarshal config error: %s", err)
	}

	for i := range p.NodeConfig {
		if p.NodeConfig[i].ApiConfig.RetryCount == nil {
			p.NodeConfig[i].ApiConfig.RetryCount = intPtr(DefaultNodeRetryCount)
		}
	}

	return nil
}

func intPtr(v int) *int {
	return &v
}
