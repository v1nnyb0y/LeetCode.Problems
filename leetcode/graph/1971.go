package graph

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int, n)
	for _, pair := range edges {
		if _, ok := graph[pair[0]]; ok {
			graph[pair[0]] = append(graph[pair[0]], pair[1])
		} else {
			graph[pair[0]] = []int{pair[1]}
		}
		if _, ok := graph[pair[1]]; ok {
			graph[pair[1]] = append(graph[pair[1]], pair[0])
		} else {
			graph[pair[1]] = []int{pair[0]}
		}
	}
	visited := make(map[int]struct{}, n)
	var dfs func(int) bool
	dfs = func(node int) bool {
		if node == destination {
			return true
		}
		visited[node] = struct{}{}
		for _, target := range graph[node] {
			if _, ok := visited[target]; !ok {
				if dfs(target) {
					return true
				}
			}
		}
		return false
	}
	return dfs(source)
}
