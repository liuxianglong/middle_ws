package code

const (
	RiskPornRateWarning  = "risk.porn_warning"
	RiskDetectionWarning = "risk.detection_warning"
	RiskBadWarning       = "risk.bad_warning"
)

var riskMap = map[string]int{
	RiskPornRateWarning:  1,
	RiskDetectionWarning: 2,
	RiskBadWarning:       3,
}
