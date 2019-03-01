package service

import (
	"model"
	"strings"
	"sync"
)

var Policy = &policy{
	mutex: &sync.Mutex{},
}

type policy struct {
	mutex *sync.Mutex
}

func (p *policy) AddPolicy(pr *model.PolicyRequest) bool {
	return enForcer.AddPolicy(pr.Role, pr.Path, strings.ToUpper(pr.Method))
}

func (p *policy) GetPolicy() [][]string {
	return enForcer.GetPolicy()
}

func (p *policy) GetAction() []string {
	return enForcer.GetAllActions()
}

func (p *policy) GetObject() []string {
	return enForcer.GetAllObjects()
}

func (p *policy) GetSubject() []string {
	return enForcer.GetAllSubjects()
}

//fieldIndex: 0 表示匹配 subject<同role>； 1 表示匹配 object<同resource>； 2 表示匹配 action<同method>
func (p *policy) DeletePolicy(fieldIndex int, value string) bool {
	//return enForcer.RemovePolicy(subject)
	return enForcer.RemoveFilteredPolicy(fieldIndex, value)
}
