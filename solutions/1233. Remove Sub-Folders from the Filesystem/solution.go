import "strings"

type Tree struct {
	value    string
	children map[string]*Tree
	isEnd    bool
}

func removeSubfolders(folder []string) []string {
	head := &Tree{
		value:    "/",
		children: make(map[string]*Tree),
		isEnd:    false,
	}

	for _, f := range folder {
		subfolders := strings.Split(strings.Trim(f, "/"), "/")
		temp := head
		for _, subfolder := range subfolders {
			if _, ok := temp.children[subfolder]; !ok {
				temp.children[subfolder] = &Tree{
					value:    subfolder,
					children: make(map[string]*Tree),
					isEnd:    false,
				}
			}

			temp = temp.children[subfolder]
		}
		temp.isEnd = true
	}

	result := make([]string, 0)
	var dfs func(node *Tree, currFolder string)
	dfs = func(node *Tree, currFolder string) {
		if node.isEnd {
			result = append(result, currFolder)
			return
		}

		for name, next := range node.children {
			nextFolder := currFolder + "/" + name
			dfs(next, nextFolder)
		}
	}

	dfs(head, "")
	return result
}
