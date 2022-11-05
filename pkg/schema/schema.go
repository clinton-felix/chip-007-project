package schema

type MetaData struct {
	Format				string		`json:"description"`
	Name				string		`json:"description"`
	Description			string		`json:"description"`
	MinitingTool		string		`json:"minting_tool"`
	SensitiveContent	bool		`json:"sensitive_content"`
	SeriesNumber		int64		`json:"series_number"`
	SeriesTotal			int64		`json:"series_total"`			
	Attributes			[]Attr1		`json:"attributes"`
	Collection			Collection	`json:"collection"`
	Data 				Data		`json:"data"`
	Hash				string		`json:"hash"`
}

type Attr1 struct{
	TraitType	string		`json:"trait_type"`
	Value		string		`json:"value"`
	MinValue	string		`json:"min_value,omitempty"`
	MaxValue	string		`json:"max_value,omitempty"`
}

type Attr2 struct{
	Type	string		`json:"type"`
	Value	string		`json:"value"`
}

type Collection struct{
	Name		string 	`json:"name"`
	ID			string	`json:"id"`
	Attributes	[]Attr2	`json:"attributes"`
}

type Data struct{
	InputData	string `json:"input_data"`
}