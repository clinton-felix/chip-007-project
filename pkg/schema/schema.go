package schema

type MetaData struct {
	Format				string		`json:"description"`
	Name				string		`json:"description"`
	Description			string		`json:"description"`
	MinitingTool		string		`json:"minting_tool"`
	SensitiveContent	bool		`json:"sensitive_content"`
	SeriesNumber		int64		`json:"series_number"`
	SeriesTotal			int64		`json:"series_total"`			
	Attributes			[]attr1		`json:"attributes"`
	Collection			collection	`json:"collection"`
	Data 				data		`json:"data"`
	Hash				string		`json:"hash"`
}

type attr1 struct{
	TraitType	string		`json:"trait_type"`
	Value		string		`json:"value"`
	MinValue	string		`json:"min_value,omitempty"`
	MaxValue	string		`json:"max_value,omitempty"`
}

type attr2 struct{
	Type	string		`json:"type"`
	Value	string		`json:"value"`
}

type collection struct{
	Name		string 	`json:"name"`
	ID			string	`json:"id"`
	Attributes	[]attr2	`json:"attributes"`
}

type data struct{
	InputData	string `json:"input_data"`
}