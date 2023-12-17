package variable

import "regexp"

var (
	Upattern, _ = regexp.Compile("^[a-zA-Z0-9]{6,20}$")
	Ppattern, _ = regexp.Compile("^[a-zA-Z0-9!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]{6,20}$")
)
