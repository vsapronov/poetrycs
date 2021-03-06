package java

import (
	"testing"
)

func TestMethodSimple(t *testing.T) {
	AssertCode(t, Method("someMethod"), `someMethod()`)
}

func TestMethodReturn(t *testing.T) {
	AssertCode(t, Method("someMethod").Returns("String"), `String someMethod()`)
}

func TestMethodWithBody(t *testing.T) {
	expected := `
someMethod() {
}
`
	method := Method("someMethod")
	method.Define().Block()
	AssertCode(t, method, expected)
}

func TestMethodPrivate(t *testing.T) {
	AssertCode(t, Method("someMethod").Private(), `private someMethod()`)
}

func TestMethodAttribute(t *testing.T) {
	AssertCode(t, Method("someMethod").Attribute("JsonCreator"), `
@JsonCreator
someMethod()`)
}
