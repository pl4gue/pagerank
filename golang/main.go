package main

import (
	"log"
	"os"

	graphs "github.com/pl4gue/pagerank/golang/src/graphs"
	pagerank "github.com/pl4gue/pagerank/golang/src/pagerank"
)

func main() {
  //Example matrix
	matrix := [][]float64{
		{0, 1, 0},
		{1, 0, 1},
		{1, 0, 0},
	}

	pagerank, err := pagerank.RankPages(matrix, 1, 0.85)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	graph := graphs.SetupGraphs("PageRank")
	nodes := graphs.GetNodesFromPageRank(pagerank)
	edges := graphs.GetLinksFromMatrix(matrix)
	graphs.AddSeriesToGraph(graph, nodes, edges)

	f, _ := os.Create("graph.html")
	graph.Render(f)
}
