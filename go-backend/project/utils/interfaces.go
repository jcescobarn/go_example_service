package utils

type Response struct {
	Format          string          `json:"format"`
	ResultsPerPage  int             `json:"resultsPerPage"`
	StartIndex      int             `json:"startIndex"`
	Timestamp       string          `json:"timestamp"`
	TotalResults    int             `json:"totalResults"`
	Version         string          `json:"version"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

type Vulnerability struct {
	Cve Cve `json:"cve"`
}

type Cve struct {
	Configurations   []Configuration `json:"configurations"`
	Descriptions     []Description   `json:"descriptions"`
	ID               string          `json:"id"`
	LastModified     string          `json:"lastModified"`
	Metrics          Metric          `json:"metrics"`
	Published        string          `json:"published"`
	References       []Reference     `json:"references"`
	SourceIdentifier string          `json:"sourceIdentifier"`
	VulnStatus       string          `json:"vulnStatus"`
	Weaknesses       []Weakness      `json:"weaknesses"`
}

type Configuration struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	CpeMatch []CpeMatch `json:"cpeMatch"`
	Negate   bool       `json:"negate"`
	Operator string     `json:"operator"`
}

type CpeMatch struct {
	Criteria        string `json:"criteria"`
	MatchCriteriaID string `json:"matchCriteriaId"`
	Vulnerable      bool   `json:"vulnerable"`
}

type Description struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type Metric struct {
	CvssMetricV2 []CvssMetricV2 `json:"cvssMetricV2"`
}

type CvssMetricV2 struct {
	AcInsufInfo             bool   `json:"acInsufInfo"`
	BaseSeverity            string `json:"baseSeverity"`
	CvssData                CvssData
	ExploitabilityScore     float64 `json:"exploitabilityScore"`
	ImpactScore             float64 `json:"impactScore"`
	ObtainAllPrivilege      bool    `json:"obtainAllPrivilege"`
	ObtainOtherPrivilege    bool    `json:"obtainOtherPrivilege"`
	ObtainUserPrivilege     bool    `json:"obtainUserPrivilege"`
	Source                  string  `json:"source"`
	Type                    string  `json:"type"`
	UserInteractionRequired bool    `json:"userInteractionRequired"`
}

type CvssData struct {
	AccessComplexity      string  `json:"accessComplexity"`
	AccessVector          string  `json:"accessVector"`
	Authentication        string  `json:"authentication"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	VectorString          string  `json:"vectorString"`
	Version               string  `json:"version"`
}

type Reference struct {
	Source string `json:"source"`
	URL    string `json:"url"`
}

type Weakness struct {
	Description []Description `json:"description"`
	Source      string        `json:"source"`
	Type        string        `json:"type"`
}

type Rule struct {
	Id            string `json:"id"`
	Last_modified string `json:"last_modified"`
	Description   string `json:"description"`
	Base_severity string `json:"severity"`
}

type RulesResponse struct {
	Total_in_page   int            `json:"total_in_page"`
	Total_pages     int            `json:"total_pages"`
	Total_results   int            `json:"total_results"`
	Index           int            `json:"index"`
	Vulnerabilities []RuleResponse `json:"vulnerabilities"`
}

type RuleResponse struct {
	Id            string `json:"id"`
	Description   string `json:"description"`
	Base_severity string `json:"severity"`
}
