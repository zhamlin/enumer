package main

import "fmt"

// Arguments to format are:
//
//	[1]: type name
const stringNameToValueMethod = `// %[1]sString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func %[1]sString(s string) (%[1]s, error) {
	if val, ok := _%[1]sNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%%s does not belong to %[1]s values", s)
}
`

// Arguments to format are:
//
//	[1]: type name
const stringValuesMethod = `// %[1]sValues returns all values of the enum
func %[1]sValues() []%[1]s {
	return _%[1]sValues
}
`

// Arguments to format are:
//
//	[1]: type name
const stringBelongsMethodLoop = `// IsA%[1]s returns "true" if the value is listed in the enum definition. "false" otherwise
func (i %[1]s) IsA%[1]s() bool {
	for _, v := range _%[1]sValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Arguments to format are:
//
//	[1]: type name
const stringBelongsMethodSet = `// IsA%[1]s returns "true" if the value is listed in the enum definition. "false" otherwise
func (i %[1]s) IsA%[1]s() bool {
	_, ok := _%[1]sMap[i]
	return ok
}
`

func (g *Generator) buildBasicExtras(runs [][]Value, typeName string, runsThreshold int) {
	// At this moment, either "g.declareIndexAndNameVars()" or "g.declareNameVars()" has been called

	// Print the slice of values
	g.Printf("\nvar _%sValues = []%s{", typeName, typeName)
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t%s, ", value.str)
		}
	}
	g.Printf("}\n\n")

	// Print the map between name and value
	g.Printf("\nvar _%sNameToValueMap = map[string]%s{\n", typeName, typeName)
	thereAreRuns := len(runs) > 1 && len(runs) <= runsThreshold
	var n int
	var runID string
	for i, values := range runs {
		if thereAreRuns {
			runID = "_" + fmt.Sprintf("%d", i)
			n = 0
		} else {
			runID = ""
		}

		for _, value := range values {
			g.Printf("\t_%sName%s[%d:%d]: %s,\n", typeName, runID, n, n+len(value.name), &value)
			n += len(value.name)
		}
	}
	g.Printf("}\n\n")

	// Print the basic extra methods
	g.Printf(stringNameToValueMethod, typeName)
	g.Printf(stringValuesMethod, typeName)

	if len(runs) <= runsThreshold {
		g.Printf(stringBelongsMethodLoop, typeName)
	} else { // There is a map of values, the code is simpler then
		g.Printf(stringBelongsMethodSet, typeName)
	}
}
