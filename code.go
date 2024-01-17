package main

type CodeOption int

type Code []CodeOption

type ComparissonResult struct {
	Correct    int
	WrongPlace int
}

func (c Code) IsValid() bool {
	codeLen := len(c)
	for i := 0; i < codeLen-1; i++ {
		for j := i + 1; j < codeLen; j++ {
			if c[i] == c[j] {
				return false
			}
		}
	}

	return true
}

func (c Code) CompareTo(otherCode Code) ComparissonResult {
	result := ComparissonResult{
		Correct:    0,
		WrongPlace: 0,
	}

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(otherCode); j++ {
			if i == j && c[i] == otherCode[j] {
				result.Correct++
			} else if c[i] == otherCode[j] {
				result.WrongPlace++
			}
		}
	}
	return result
}
