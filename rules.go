package main

import "fmt"

type Rules struct {
	Rules map[string]string
}

func NewRules() *Rules {
	instance := &Rules{
		Rules: make(map[string]string),
	}
	instance.InitRules()
	return instance
}

func (me *Rules) AddRule(name, regex string) {
	me.Rules[name] = regex
}

func (me *Rules) AddRuleMap(rule map[string]string) {
	for ruleKey, ruleVal := range rule {
		me.Rules[ruleKey] = ruleVal
	}
}

func (me *Rules) InitRules() {
	me.Rules["email"] = `[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`
	me.Rules["domain"] = `(?m)(?:https?://)?(?:www\.)?([a-zA-Z0-9-]+(?:\.[a-zA-Z]+)+)`
}

func (me *Rules) PrintAvailables() {
	fmt.Println("Rules availables")
	for ruleKey, ruleVal := range me.Rules {
		fmt.Printf("> %q = %q \n", ruleKey, ruleVal)
	}
	fmt.Println("")
}
