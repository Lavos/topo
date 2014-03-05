package main

import (
	"log"
	"bufio"
	"os"
)

type Node struct {
	Value rune
	Next *Node
}

type List struct {
	First *Node
	Nodes map[rune]*Node
}

func (l *List) Insert(r rune) (*Node, bool) {
	// check if the node exists for this character
	if _, ok := l.Nodes[r]; ok {
		return l.Nodes[r], false
	}

	n := &Node{
		Value: r,
	}

	l.Nodes[r] = n

	return n, true
}

func (n *Node) InsertAfter(newNode *Node) {
	newNode.Next = n.Next
	n.Next = newNode
}

func main () {
	list := &List{
		First: &Node{ Value: '0', },
		Nodes: make(map[rune]*Node),
	}

	list.Nodes['0'] = list.First

	var current_word []byte
	previous_word := []byte{'0'}
	var a, b rune
	var a_node, b_node *Node
	var b_new_node bool

	// create nodes based on list
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		current_word = scanner.Bytes()

		log.Printf("comparing previous_word: %q to current_word: %q", previous_word, current_word)

		for index, char := range previous_word {
			log.Printf("index: %v, char: %q:", index, char)

			if len(current_word)-1 < index || char == '\n' {
				log.Print("break early 1")
				break
			}

			a = rune(char)
			b = rune(current_word[index])

			a_node, _ = list.Insert(a)
			b_node, b_new_node = list.Insert(b)

			if a != b && b_new_node {
				log.Printf("LOGIC %q comes before %q", a, b)
				a_node.InsertAfter(b_node)
				break
			}
		}

		previous_word = current_word
	}

	log.Printf("%#v", list)
	log.Printf("list has %#v nodes", len(list.Nodes))

	// iterate over list

	for _, n := range list.Nodes {
		if n.Next != nil {
			log.Printf("%q Next %q", n.Value, n.Next.Value)
		} else {
			log.Printf("%q Next nil", n.Value)
		}
	}

	/* node := list.First
	for node != nil {
		log.Printf("%q", node.Value)
		node = node.Next
	} */
}
