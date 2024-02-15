package dfa

func parse(src string) (out *Config, err error) {
	return &Config{"start", ConfigState{
		"start": {
			"[1-9]": {"int"},
			"0":     {"zer"},
			"-":     {"neg"},
		},
		"int": {
			"eof":   {"success"},
			"[0-9]": {"int"},
			"[eE]":  {"sci"},
			".":     {"dot"},
		},
		"neg": {
			"eof":   {"success"},
			"[1-9]": {"int"},
			"0":     {"zer"},
		},
		"dig": {
			"eof":   {"success"},
			"[0-9]": {"dig"},
			"[eE]":  {"sci"},
			".":     {"dot"},
		},
		"zer": {
			"eof":  {"success"},
			"[eE]": {"sci"},
			".":    {"dot"},
		},
		"dot": {
			"[0-9]": {"dec"},
		},
		"dec": {
			"eof":   {"success"},
			"[0-9]": {"dec"},
			"[eE]":  {"sci"},
		},
		"sci": {
			"[0-9]": {"exp"},
			"[+-]":  {"sgn"},
		},
		"sgn": {
			"[0-9]": {"exp"},
		},
		"exp": {
			"eof":   {"success"},
			"[0-9]": {"exp"},
		},
	}}, nil
}
