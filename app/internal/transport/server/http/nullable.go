package http

type NullableString struct {
    Value string
    Set   bool
    Null  bool
}

func (s *NullableString) UnmarshalJSON(text []byte) error {
    s.Set = true
    s.Null = false

    if string(text) == "" {
        s.Null = true
        return nil
    }

    s.Value = string(text)
    return nil
}
