package nullable

type String struct {
	Value string
	Set   bool
	Null  bool
}

func (s *String) UnmarshalJSON(text []byte) error {
	s.Set = true
	s.Null = false

	if string(text) == "" {
		s.Null = true
		return nil
	}

	s.Value = string(text)
	return nil
}
