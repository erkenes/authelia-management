package configuration

import (
	"fmt"
	"github.com/imdario/mergo"
)

func HasAccessRuleForDomain(domain string) bool {
	return hasAccessRuleForDomain(domain)
}

func AddAccessRuleForDomain(domain []string, policy string, subjects []string, resources []string) {
	configuration := getConfiguration()

	var newRule AutheliaConfigStruct
	newRule = addAccessRuleForDomain(domain, policy, subjects, resources)

	mergo.Merge(&configuration, newRule)

	fmt.Println(configuration)
	fmt.Println(newRule)

	saveToFile(configuration)
}
