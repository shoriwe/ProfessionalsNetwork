package ArgumentParser

import (
	"fmt"
	"os"
)

func readEnvironmentVariable(variableName string) interface{} {
	value, found := os.LookupEnv(variableName)
	if found {
		_ = os.Setenv(variableName, "NONE")
		return value
	}
	return nil
}

func GetEnvironmentVariables() map[string]string {
	variables := map[string]string{}
	for _, variable := range EnvironmentVariables {
		value := readEnvironmentVariable(variable)
		if value != nil {
			if len(value.(string)) > 0 {
				variables[variable] = value.(string)
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "Variable \"%s\" can't be empty", variable)
				os.Exit(-1)
			}
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "Variable \"%s\" required", variable)
			os.Exit(-1)
		}
	}
	return variables
}

