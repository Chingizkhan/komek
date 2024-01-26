package word

const (
	English Language = "en"
	Kazakh  Language = "kz"
)

var languages = []Language{
	English,
	Kazakh,
}

type (
	Language string
)

func (l *Language) Contains() bool {
	var exists bool
	for _, ls := range languages {
		if ls == *l {
			exists = true
		}
	}
	return exists
}

func (l *Language) Validate() error {
	if !l.Contains() {
		return ErrIncorrectLanguage
	}
	return nil
}

type Word struct {
	Value       string
	Language    Language
	Translation string
}

func (w *Word) Validate() error {
	if err := w.Language.Validate(); err != nil {
		return err
	}
	return nil
}
