package generator

type ruler struct {
	// all codes generated are unique (no duplicate codes)
	isUniqueCode bool
	// no duplicate character per one code
	isUniqueCharacter bool

	prefix string

	length int
}

func defaultRuler() *ruler {
	return &ruler{
		isUniqueCode:      true,
		isUniqueCharacter: true,
		prefix:            "CR",
		length:            8,
	}
}
