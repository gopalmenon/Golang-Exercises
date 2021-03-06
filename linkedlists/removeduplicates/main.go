package removeduplicates

import "golangexercises/linkedlists"

func removeDupsWithHashing(l *linkedlists.List) {

	if l == nil || l.Head == nil {
		return
	}

	hashtable := make(map[interface{}]int)
	for node := l.Head; node != nil; node = node.Next {
		if _, ok := hashtable[node.Value]; ok {
			l.Remove(node)
		} else {
			hashtable[node.Value] += 1
		}
	}

}

func removeDupsWithRunners(l *linkedlists.List) {

	if l == nil || l.Head == nil {
		return
	}

	runner1 := l.Head
	if runner1.Next == nil {
		return
	}

	for ; runner1 != nil; runner1 = runner1.Next {

		runner2 := runner1.Next
		if runner2 == nil {
			return
		}

		for runner2 != nil {

			nextRunner := runner2.Next
			if runner2.Value == runner1.Value {
				l.Remove(runner2)
			}
			runner2 = nextRunner
		}

	}

}
