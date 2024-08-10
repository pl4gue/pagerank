# PageRank Algorithm

The PageRank algorithm is a Google algorithm that ranks web pages in search
results by evaluating the number and quality of links to a page. It operates on
the principle that pages receiving more high-quality links are deemed more
important and are thus ranked higher.

And this is my implementation of it, in various languages.

> [!note]
> It's not very impressive or well tought or anything, but it's fun to try and implement it since it
> deals with matrices, arrays, and you can try different approaches each time.

## Explanation of the overall implementation

Given a square array of size $n$ where each column $i$ and row $j$ represents the presence of a
hyperlink leading from the page $i$ to the page $j$.

We calculate the pagerank of each page over various iterations of the following algorithm.

```python
# Node count
n = len(matrix)

# Initialize pagerank as an uniform vector which total sum is 1
pagerank = [1 / n] * n

# Iterates over and over
for _ in range(iterations):
    new_pagerank = [(1 - damping_factor) / n] * n

    for i in range(n):
        for j in range(n):
            # Skips links to itself
            if j == i:
                continue

            # adj_matrix[j][i]: (0 or 1) Whether there's a link outgoing from page j to page i.
            # average_or_zero:  The average between page j's rank and the amount of outgoing
            #                   links from it, if there's no outgoing link a division by zero
            #                   might occur so just set it's result by 0 because the previous
            #                   value will be 0.
            # damping_factor:   A damping factor to account for randomness from users, usually 0.85.
            new_pagerank[i] += (
                adj_matrix[j][i]
                * average_or_zero(pagerank[j], sum(adj_matrix[j]))
                * damping_factor
            )

    pagerank = new_pagerank
```

## Current implementations

- [x] Python script
  - Graph visualization using [pyvis](https://github.com/WestHealth/pyvis)
  - Run by using
    ```sh
    cd python/
    pip install -r requirements.txt
    python page_rank.py
    ```
- [x] Golang script
  - Graph visualization using [go-echarts](https://github.com/go-echarts/go-echarts)
  - Run by using
    ```sh
    cd golang/
    go build
    ./main
    # or
    go run .
    ```
