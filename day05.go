package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"time"
)

func Day05Part1(scanner *bufio.Scanner) int {
	type instruction struct {
		ParentID int
		ChildID  int
		Update   []int
	}

	var instructions []*instruction
	instructions = parse(scanner, func(line string) *instruction {
		if line == "" {
			return nil
		}
		i := &instruction{}
		order := strings.Split(line, "|")
		if len(order) == 2 {
			i.ParentID = mustInt(order[0])
			i.ChildID = mustInt(order[1])
			return i
		}
		for _, num := range strings.Split(line, ",") {
			i.Update = append(i.Update, mustInt(num))
		}
		return i
	})

	type Node struct {
		ID       int
		Parents  []*Node
		Children []*Node
		Weight   int
	}
	type Graph struct {
		nodes map[int]*Node
	}

	g := Graph{
		nodes: map[int]*Node{},
	}
	getNode := func(nodeID int) *Node {
		n := g.nodes[nodeID]
		if n != nil {
			return n
		}
		n = &Node{ID: nodeID}
		g.nodes[nodeID] = n
		return n
	}
	dotGraph := func() {
		fmt.Println("digraph {")
		for _, node := range g.nodes {
			fmt.Printf("  %d [label=%d, xlabel=%d];\n", node.ID, node.ID, node.Weight)
			for _, parent := range node.Parents {
				fmt.Printf("  %d -> %d;\n", parent.ID, node.ID)
			}
			// 	for _, child := range node.Children {
			// 		fmt.Printf("  %d -> %d;\n", node.ID, child.ID)
			// 	}
		}
		fmt.Println("}")
	}
	findRoot := func() *Node {
		var loc *Node
		for _, node := range g.nodes {
			loc = node
			break
		}
		for {
			if len(loc.Parents) == 0 {
				return loc
			}
			loc = loc.Parents[0]
		}
	}

	// Load all nodes into the graph
	for _, i := range instructions {
		if i.ParentID == 0 {
			continue
		}
		parent := getNode(i.ParentID)
		child := getNode(i.ChildID)
		parent.Children = append(parent.Children, child)
		child.Parents = append(child.Parents, parent)
	}
	_ = findRoot
	_ = dotGraph
	// node := findRoot()
	// node.Weight = 1
	// fmt.Println(node)
	// dotGraph()

	var total int
	for _, i := range instructions {
		if i.ParentID > 0 {
			continue
		}
		valid := true
		for idx := 0; idx < len(i.Update)-1; idx++ {
			nodeID := i.Update[idx]
			node := g.nodes[nodeID]

			var found bool
			for _, child := range node.Children {
				if child.ID == i.Update[idx+1] {
					found = true
				}
			}
			if !found {
				valid = false
			}
		}
		if valid {
			fmt.Println("valid", i.Update[len(i.Update)/2], i.Update)
			total += i.Update[len(i.Update)/2]
		} else {
			fmt.Println("invalid", i.Update)
		}
	}

	return total
}

func Day05Part2(scanner *bufio.Scanner) int {
	type instruction struct {
		ParentID int
		ChildID  int
		Update   []int
	}

	var instructions []*instruction
	instructions = parse(scanner, func(line string) *instruction {
		if line == "" {
			return nil
		}
		i := &instruction{}
		order := strings.Split(line, "|")
		if len(order) == 2 {
			i.ParentID = mustInt(order[0])
			i.ChildID = mustInt(order[1])
			return i
		}
		for _, num := range strings.Split(line, ",") {
			i.Update = append(i.Update, mustInt(num))
		}
		return i
	})

	type Node struct {
		ID       int
		Parents  []*Node
		Children []*Node
		Weight   int
	}
	type Graph struct {
		nodes map[int]*Node
	}

	g := Graph{
		nodes: map[int]*Node{},
	}
	getNode := func(nodeID int) *Node {
		n := g.nodes[nodeID]
		if n != nil {
			return n
		}
		n = &Node{ID: nodeID}
		g.nodes[nodeID] = n
		return n
	}
	dotGraph := func() {
		fmt.Println("digraph {")
		for _, node := range g.nodes {
			fmt.Printf("  %d [label=%d, xlabel=%d];\n", node.ID, node.ID, node.Weight)
			for _, parent := range node.Parents {
				fmt.Printf("  %d -> %d;\n", parent.ID, node.ID)
			}
			// 	for _, child := range node.Children {
			// 		fmt.Printf("  %d -> %d;\n", node.ID, child.ID)
			// 	}
		}
		fmt.Println("}")
	}
	findRoot := func() *Node {
		var loc *Node
		for _, node := range g.nodes {
			loc = node
			break
		}
		for {
			if len(loc.Parents) == 0 {
				return loc
			}
			loc = loc.Parents[0]
		}
	}

	// Load all nodes into the graph
	for _, i := range instructions {
		if i.ParentID == 0 {
			continue
		}
		parent := getNode(i.ParentID)
		child := getNode(i.ChildID)
		parent.Children = append(parent.Children, child)
		child.Parents = append(child.Parents, parent)
	}
	_ = findRoot
	_ = dotGraph
	// node := findRoot()
	// node.Weight = 1
	// fmt.Println(node)
	// dotGraph()

	fixInvalid := func(update []int) []int {
		fixed := slices.Clone(update)
		for {
			done := true
			for i := 0; i < len(fixed)-1; i++ {
				nodeID := fixed[i]
				node := g.nodes[nodeID]

				var found bool
				for _, child := range node.Children {
					if child.ID == update[i+1] {
						found = true
					}
				}
				if !found {
					fixed[i], fixed[i+1] = fixed[i+1], fixed[i]
					done = false
				}
			}
			fmt.Println(fixed)
			time.Sleep(3 * time.Second)
			if done {
				break
			}
		}
		return fixed
	}

	var total int
	for _, i := range instructions {
		if i.ParentID > 0 {
			continue
		}
		valid := true
		for idx := 0; idx < len(i.Update)-1; idx++ {
			nodeID := i.Update[idx]
			node := g.nodes[nodeID]

			var found bool
			for _, child := range node.Children {
				if child.ID == i.Update[idx+1] {
					found = true
				}
			}
			if !found {
				valid = false
			}
		}
		if valid {
			fmt.Println("valid", i.Update[len(i.Update)/2], i.Update)
		} else {
			fmt.Println("invalid", i.Update)
			fixed := fixInvalid(i.Update)
			fmt.Println("fixed", fixed)
			total += fixed[len(fixed)/2]
		}
	}

	return total
}
