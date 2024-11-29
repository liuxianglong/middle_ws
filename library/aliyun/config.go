package aliyun

// KeyConfig 定义了阿里云各个服务需要的参数
type KeyConfig struct {
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
	Secret   string `json:"secret"`
}

type KeyLocalConfig struct {
	Endpoint     string `json:"endpoint"`
	Bucket       string `json:"bucket"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	CdnDomain    string `json:"cdnDomain"`
}

//func parseToConfigPt(content string, pt *KeyConfig) error {
//	v := &KeyConfig{}
//	err := json.Unmarshal([]byte(content), v)
//	if err != nil {
//		return err
//	}
//	*pt = *v
//	return nil
//}
