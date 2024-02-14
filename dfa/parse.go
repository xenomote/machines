package dfa

func parse(src string) (out *Config, err error) {
	return &Config{"start", ConfigState{
		"start": {
			"[1-9]": {"int", "output"},
			"0":     {"zer"},
			"-":     {"neg"},
		},
		"int": {
			"[0-9]": {"int"},
			"[eE]":  {"sci"},
			".":     {"dot"},
			"eof":   {"success"},
		},
		"neg": {
			"[1-9]": {"int"},
			"0":     {"zer"},
			"eof":   {"success"},
		},
		"dig": {
			"[0-9]": {"dig"},
			"[eE]":  {"sci"},
			".":     {"dot"},
			"eof":   {"success"},
		},
		"zer": {
			"[eE]": {"sci"},
			".":    {"dot"},
			"eof":  {"success"},
		},
		"dot": {
			"[0-9]": {"dec"},
		},
		"dec": {
			"[0-9]": {"dec"},
			"[eE]":  {"sci"},
			"eof":   {"success"},
		},
		"sci": {
			"[0-9]": {"exp"},
			"[+-]":  {"sgn"},
		},
		"sgn": {
			"[0-9]": {"exp"},
		},
		"exp": {
			"[0-9]": {"exp"},
			"eof":   {"success"},
		},
	}}, nil
}
