package common

import "github.com/geneva/when/rules"

var All = []rules.Rule{
	SlashDMY(rules.Override),
}
