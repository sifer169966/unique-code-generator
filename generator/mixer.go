package generator

type mixer struct {
	upperCaseEN string
	numeric     string
}

type mixerOption func(*mixer)

func WithOptionUpperCaseEN(upperCaseEN string) mixerOption {
	return func(m *mixer) {
		m.upperCaseEN = upperCaseEN
	}
}

func WithOptionNumberic(numberic string) mixerOption {
	return func(m *mixer) {
		m.numeric = numberic
	}
}

func defaultMixer() *mixer {
	const defaultUpperCaseEN = "ABCDEFGHIGKLMNOPQRSTUVWXYZ"
	const defaultNumeric = "1234567890"
	return &mixer{
		upperCaseEN: defaultUpperCaseEN,
		numeric:     defaultNumeric,
	}
}

func NewMixer(mixerOptions ...mixerOption) *mixer {
	mixer := defaultMixer()
	for _, opt := range mixerOptions {
		opt(mixer)
	}
	return mixer
}
