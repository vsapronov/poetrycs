package scala

import (
	"strings"
)

type TraitDeclaration struct {
	name       string
	extends    []string
	modifiers  []string
	attributes []Writable
	members    *StatementsDeclaration
}

func (self *TraitDeclaration) addModifier(modifier string) *TraitDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *TraitDeclaration) Private() *TraitDeclaration {
	return self.addModifier("private")
}

func (self *TraitDeclaration) Public() *TraitDeclaration {
	return self.addModifier("public")
}

func (self *TraitDeclaration) Sealed() *TraitDeclaration {
	return self.addModifier("sealed")
}

func (self *TraitDeclaration) Extends(types ...string) *TraitDeclaration {
	self.extends = append(self.extends, types...)
	return self
}

func (self *TraitDeclaration) Members(members ...Writable) *TraitDeclaration {
	self.members = Scope(members...)
	return self
}

func (self *TraitDeclaration) MembersInline(members ...Writable) *TraitDeclaration {
	self.members = ScopeInline(members...)
	return self
}

func (self *TraitDeclaration) Add(members ...Writable) *TraitDeclaration {
	if self.members == nil {
		self.members = Scope()
	}
	self.members.Add(members...)
	return self
}

func (self *TraitDeclaration) AddAttributes(attributes ...Writable) *TraitDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *TraitDeclaration) Attribute(code string) *TraitDeclaration {
	return self.AddAttributes(Attribute(code))
}

func Trait(name string) *TraitDeclaration {
	return &TraitDeclaration{
		name:       name,
		attributes: []Writable{},
		members:    nil,
	}
}

func (self *TraitDeclaration) WriteCode(writer CodeWriter) {
	for _, attribute := range self.attributes {
		attribute.WriteCode(writer)
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}

	writer.Write("trait ")
	writer.Write(self.name)

	if len(self.extends) > 0 {
		writer.Write(" extends " + strings.Join(self.extends, " with "))
	}

	if self.members != nil {
		writer.Write(" ")
		self.members.WriteCode(writer)
	} else {
		writer.Eol()
	}
}
