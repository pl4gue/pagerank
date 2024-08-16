package pagerank

import (
	"errors"
)

/*
Calculates the rank of each page on the given adjacency matrix of type T.

T is any real number.

Parameters:

  adj_matrix [][]T:   Adjacency matrix, always a square matrix.
  iter         int:   Algorithm iteration count. Default = 1
  damp_factor    T:   Damping factor. Default = 0.85.

Returns:
	[]T:                Result list of page ranks.
	error:              Returns in case the matrix lenght is 0
*/

func RankPages[T Real](adj_matrix [][]T, iter int, damp_factor T) ([]T, error) {
	// Node count
	n := len(adj_matrix)

	if n == 0 {
		return nil, errors.New("Matrix has size 0")
	}

	// Normalizes the value.
	// Returns 0 when dividing by zero
	average_or_zero := Divide[T]

	t_n := T(n)

	// Uniform vector inicialization
	pagerank := generate_uniform_vector(n, 1/t_n)

	// PageRank algorithm
	for range iter {
		new_pagerank := generate_uniform_vector(n, (1-damp_factor)/t_n)

		for i := range n {
			for j := range n {
				if i == j {
					continue
				}

				new_pagerank[i] += adj_matrix[j][i] *
					average_or_zero(pagerank[j], Sum(adj_matrix[j]...)) *
					damp_factor
			}
		}

		pagerank = new_pagerank
	}

	return pagerank, nil
}

// Returns an uniform vector of size n, and applies the same result of type T to all elements.
func generate_uniform_vector[T any](n int, v T) []T {
	result_matrix := make([]T, n)

	for i := range result_matrix {
		result_matrix[i] = v
	}

	return result_matrix
}
