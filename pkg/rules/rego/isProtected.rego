package protectionRules

import repo.v1

default rule := false

rule if input.is_protected == false