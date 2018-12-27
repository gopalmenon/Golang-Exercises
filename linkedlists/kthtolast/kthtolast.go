package kthtolast

import "golangexercises/linkedlists"

func kthToLastElement(l *linkedlists.List, k int) *linkedlists.Node {

	kLongList := linkedlists.New()

	for nextNode, counter := l.Head, 0; nextNode != nil; nextNode, counter = nextNode.Next, counter+1 {

		kLongList.Pushback(nextNode.Value)
		if counter > k {
			kLongList.Remove(kLongList.Head)
		}

	}

	return kLongList.Head

}
