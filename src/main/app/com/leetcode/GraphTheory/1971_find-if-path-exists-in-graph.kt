package com.leetcode.GraphTheory

class `1971_find-if-path-exists-in-graph` {
    var result = false

    fun validPath(n: Int, edges: Array<IntArray>, start: Int, end: Int): Boolean {
        if (start == end) return true

        val graph = mutableMapOf<Int, MutableList<Int>>()
        val visited = BooleanArray(n)

        for (vertix in 0 until n) {
            graph[vertix] = mutableListOf<Int>()
        }

        for (edge in edges) {
            val accessStart: MutableList<Int> = graph[edge[0]] ?: mutableListOf()
            accessStart.add(edge[1])
            graph[edge[0]] = accessStart

            val accessEnd: MutableList<Int> = graph[edge[1]] ?: mutableListOf()
            accessEnd.add(edge[0])
            graph[edge[1]] = accessEnd
        }

        findAWayDFS(
            graph,
            start,
            end,
            visited
        )

        return result
    }

    fun findAWayDFS(
        access: MutableMap<Int, MutableList<Int>>,
        start: Int,
        end: Int,
        visited: BooleanArray
    ) {
        if (visited[start]) return

        visited[start] = true

        for (vertix in access[start] ?: mutableListOf()) {
            if (vertix == end) {
                result = true
                return
            } else {
                findAWayDFS(access, vertix, end, visited)
            }
        }
    }
}
