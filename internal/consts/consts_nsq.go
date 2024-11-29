package consts

const (
	NsqConfigLookupName = "go-nsq.lookup"
	//NsqGroupDefault 默认的NSD机器组配置名字
	NsqGroupDefault = "default"
	//NsqConfigNsdName 在config中对应的配置名字
	NsqConfigNsdName = "go-nsq"
	TopicDemo        = "demo"
)

var (
	Topics = map[string]string{
		TopicDemo: NsqGroupDefault,
	}
)
