package main

import "fmt"

func isPersonSeller(name string) bool {
	// Глупая проверка: если имя человека заканчивается на "m", то он продаёт
	// манго
	return string(name[len(name)-1]) == "m"
}

func bfs(aGraph map[string][]string, name string) bool {
	var searchQueue []string
	searchQueue = append(searchQueue, aGraph[name]...)
	var searched []string

	for len(searchQueue) > 0 {
		var person = searchQueue[0]
		searchQueue = searchQueue[1:]
		isPersonAlreadySearched := false

		for i := 0; i < len(searched); i++ {
			if searched[i] == person {
				isPersonAlreadySearched = true
			}
		}

		if isPersonAlreadySearched == false {
			if isPersonSeller(person) {
				fmt.Println(person, "is a mango seller!")
				return true
			}

			searchQueue = append(searchQueue, aGraph[person]...)
			searched = append(searched, person)
		}
	}
	return false
}

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	bfs(graph, "you")
}
