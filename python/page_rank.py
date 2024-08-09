from pyvis.network import Network
from random import choice

def page_rank(adj_matrix: list[list[int]], iter: int = 1, damp_factor: float = 0.85) -> list[float]:
    """
    Calculates the rank of each page on the given adjacency matrix.

    Parameters:
        matrix (list[list[int]]): Adjacency matrix.
        iter (int):               Algorithm iteration count. Default = 1
        d (float):                Damping factor. Default = 0.85.

    Returns:
        list[float]:              Result list of page ranks.
    """

    # Node count
    n = len(adj_matrix)

    # Initialize pagerank as an uniform vector
    pagerank = [1 / n] * n

    # PageRank algorithm
    for _ in range(iter):
        new_pagerank = [(1 - damp_factor) / n] * n

        for i in range(n):
            for j in range(n):
                if j == i:
                    continue

                new_pagerank[i] += (
                    damp_factor * normal_value(adj_matrix[j][i], sum(adj_matrix[j])) * pagerank[j]
                )

        pagerank = new_pagerank

    return pagerank

# Divisions by zero return 0 in this algorithm
def normal_value(m: int, s: int) -> float:
    return 0 if s == 0 else m / s

# Generates a random matrix
def generate_matrix(n: int) -> list[list[int]]:
    return [[choice([0, 0, 1]) for _ in range(n)] for _ in range(n)]


if __name__ == "__main__":
    net = Network(directed=True)

    # Example matrix
    # matrix = [
    #     [0, 1, 0],
    #     [1, 0, 1],
    #     [1, 0, 0],
    #]

    matrix = generate_matrix(5)

    n = len(matrix)
    pagerank = page_rank(matrix)

    net.add_nodes(
        [i for i in range(n)],
        value=pagerank,
        label=[chr(i + 65) for i in range(n)],
    )

    for i in range(n):
        for j in range(n):
            if matrix[i][j] == 1:
                net.add_edge(i, j, arrows="to")

    net.show("graph.html", notebook=False)
