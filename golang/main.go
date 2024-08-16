package main

import (
	"log"
	"os"

	pagerank "github.com/pl4gue/pagerank/golang/internal"
)

func main() {
  //Example matrix
	matrix := [][]float64{
		{0, 1, 0},
		{1, 0, 1},
		{1, 0, 0},
	}

	page_rank, err := pagerank.RankPages(matrix, 1, 0.85)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	graph := pagerank.SetupGraphs("PageRank")
	nodes := pagerank.GetNodesFromPageRank(page_rank)
	edges := pagerank.GetLinksFromMatrix(matrix)
	pagerank.AddSeriesToGraph(graph, nodes, edges)

	f, _ := os.Create("graph.html")
	graph.Render(f)
}
