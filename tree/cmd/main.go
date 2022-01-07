package main

import (
	"fmt"
	"struct_tree_and_que/tree/model"
)

var (
	Ivan = model.Human{
		FirstName: "Ivan",
		LastName:  "Konev",
		Age:       21,
		Iq:        0,
	}
	Elsa = model.Human{
		FirstName: "Elsa",
		LastName:  "Smith",
		Age:       23,
		Iq:        119,
	}
	John = model.Human{
		FirstName: "John",
		LastName:  "Davidson",
		Age:       32,
		Iq:        144,
	}
	Diana = model.Human{
		FirstName: "Diana",
		LastName:  "Harper",
		Age:       7,
		Iq:        89,
	}
	Igor = model.Human{
		FirstName: "Igor",
		LastName:  "Petrov",
		Age:       14,
		Iq:        99,
	}
	Oleg = model.Human{
		FirstName: "Oleg",
		LastName:  "Frolov",
		Age:       46,
		Iq:        152,
	}
	Dima = model.Human{
		FirstName: "Dima",
		LastName:  "Onechko",
		Age:       28,
		Iq:        112,
	}
)

func main() {
	var t model.Tree
	t.Insert(Elsa)
	t.Insert(Igor)
	t.Insert(Ivan)
	t.Insert(John)
	t.Insert(Oleg)
	t.Insert(Dima)
	t.Insert(Diana)

	key := Elsa
	node, err := t.Find(key)
	if err != nil {
		fmt.Printf("error finding node: %v", err)
		return
	}
	fmt.Printf("finded node: %#v\n", node)

	fmt.Printf("remove key: %v\n", key)

	fmt.Println("Before:")
	printTree(t.Root, 0, "ROOT")

	if err := t.Remove(key); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("After: ")
	printTree(t.Root, 0, "ROOT")
}

func printTree(n *model.Node, ns int, ch string) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s:%v\n", ch, n.Key)
	printTree(n.Left, ns+2, "L")
	printTree(n.Right, ns+2, "R")
}
