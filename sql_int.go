package main

// Arguments to format are:
//	[1]: type name
const intValueMethod = `func (i %[1]s) Value() (driver.Value, error) {
	return int(i), nil
}
`

const intScanMethod = `func (i *%[1]s) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	intVal, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

    t := %[1]s(intVal)
    if !t.IsA%[1]s() {
        return fmt.Errorf("%%d does not beling to %[1]s", intVal)
    }

	*i = t
	return nil
}
`

func (g *Generator) addIntValueAndScanMethod(typeName string) {
	g.Printf("\n")
	g.Printf(intValueMethod, typeName)
	g.Printf("\n\n")
	g.Printf(intScanMethod, typeName)
}
