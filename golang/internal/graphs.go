package pagerank

import (
  "strconv"

	charts "github.com/go-echarts/go-echarts/v2/charts"
	opts "github.com/go-echarts/go-echarts/v2/opts"
)

func SetupGraphs(title string) *charts.Graph {
	graph := charts.NewGraph()

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
  )

  return graph
}

func AddSeriesToGraph(graph *charts.Graph, nodes []opts.GraphNode, edges []opts.GraphLink) {
  graph.AddSeries(graph.Title.Title, nodes, edges).
    SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Layout:             "circular",
				Roam:               opts.Bool(true),
				FocusNodeAdjacency: opts.Bool(true),
				Draggable:          opts.Bool(true),
			}),
			charts.WithLabelOpts(opts.Label{
				Show:        opts.Bool(true),
				Position:    "bottom",
				FontStyle:   "normal",
				BorderWidth: 0,
			}),
			charts.WithLineStyleOpts(opts.LineStyle{
				Curveness: .3,
			}),
    )
}

func GetNodesFromPageRank(pagerank []float64) []opts.GraphNode {
  n := len(pagerank)
  
	nodes := make([]opts.GraphNode, n)

	for i := range nodes {
		nodes[i] = opts.GraphNode{
			Name:       strconv.Itoa(i),
			Value:      float32(pagerank[i]),
			SymbolSize: pagerank[i] * 100,
		}
	}

  return nodes
}

func GetLinksFromMatrix(matrix [][]float64) []opts.GraphLink {
	var edges []opts.GraphLink

  n := len(matrix)

	for i := range n {
		for j := range n {
			if matrix[i][j] == 1 {
				edges = append(edges, opts.GraphLink{
					Source: strconv.Itoa(i),
					Target: strconv.Itoa(j),
				})
			}
		}
	}

  return edges
}
