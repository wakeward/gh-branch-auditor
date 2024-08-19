package eval

type Reports struct {
	Title  string   `json:"title"`
	Verson string   `json:"version"`
	Date   string   `json:"date"`
	Report []Report `json:"report"`
}

type Report struct {
	Repo   string    `json:"repo"`
	Branch string    `json:"branch"`
	Rules  []RuleRef `json:"rules"`
}

type RuleRef struct {
	ID              string `json:"id"`
	Rule            string `json:"rule"`
	Risk            string `json:"risk"`
	Severity        string `json:"severity"`
	ProtectionRules int    `json:"-"`
}
