package main

type Name string

func (n Name) Change() {
}

type Unvalid *string

func (u Unvalid) A() { // compile error: invalid receiver Unvalid (pointer or interface type)compiler
}

func main() {

}
