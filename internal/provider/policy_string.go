package provider

import (
	"fmt"
	"strings"
)

func (d *Directive) GenerateDirective() string {
	var directiveStrings []string

	directiveStrings = append(directiveStrings, d.Name.ValueString())

	for _, keyword := range d.Keywords {
		directiveStrings = append(directiveStrings, fmt.Sprintf("'%s'", keyword.ValueString()))
	}

	for _, scheme := range d.Schemes {
		directiveStrings = append(directiveStrings, fmt.Sprintf("%s:", scheme.ValueString()))
	}

	for _, host := range d.Hosts {
		directiveStrings = append(directiveStrings, host.ValueString())
	}

	for _, nonce := range d.Nonces {
		directiveStrings = append(directiveStrings, fmt.Sprintf("'nonce-%s'", nonce.ValueString()))
	}

	for _, hash := range d.Hashes {
		directiveStrings = append(directiveStrings, fmt.Sprintf("'%s-%s'", hash.Algorithm.ValueString(), hash.Value.ValueString()))
	}

	for _, value := range d.Values {
		directiveStrings = append(directiveStrings, value.ValueString())
	}

	return fmt.Sprintf("%s;", strings.Join(directiveStrings, " "))
}

func (p *PolicyDataSourceModel) GeneratePolicy() string {
	directiveStrings := make([]string, len(p.Directives))
	for i, directive := range p.Directives {
		directiveStrings[i] = directive.GenerateDirective()
	}

	return strings.Join(directiveStrings, " ")
}
