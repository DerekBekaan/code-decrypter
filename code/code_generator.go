package code

type LinkedCodeOption struct {
	option         CodeOption
	previousOption *LinkedCodeOption
}

func (l *LinkedCodeOption) next() bool {
	var numberOfOptions CodeOption = 7
	if l.option < numberOfOptions {
		l.option++
	} else if l.previousOption != nil {
		l.option = 0
		return l.previousOption.next()
	} else {
		return false
	}

	return true
}

func (l *LinkedCodeOption) extractCode(code *Code) {
	*code = append(*code, l.option)

	if l.previousOption != nil {
		l.previousOption.extractCode(code)
	}
}

func GenerateAllCodes() []Code {
	allCodes := make([]Code, 6720)

	optionsInCode := 5
	code := make([]LinkedCodeOption, optionsInCode)

	// initialize code to 0s and link them
	for i := 0; i < optionsInCode; i++ {
		var previousOption *LinkedCodeOption = nil
		if i > 0 {
			previousOption = &code[i-1]
		}

		code[i] = LinkedCodeOption{
			option:         0,
			previousOption: previousOption,
		}
	}

	count := 0
	stillHaveCodes := true
	for stillHaveCodes {
		var extractedCode Code
		code[optionsInCode-1].extractCode(&extractedCode)

		if extractedCode.IsValid() {
			allCodes[count] = extractedCode
			count++
		}

		stillHaveCodes = code[optionsInCode-1].next()
	}

	return allCodes
}
