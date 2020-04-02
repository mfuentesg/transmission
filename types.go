package transmission

type NumBool bool

func (bit *NumBool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = txt == "1" || txt == "true"

	return nil
}
