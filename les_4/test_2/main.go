package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

/*
	  Идея: первое значение в списке получает ссылку на следующее значение,
	последнее значение получает ссылку на nil.
	  Идея: для удаления дупликатов необходимо пройтись по списку с проверкой
	следующего значения, если они совпадают, то присвоить ссылку следующего
	значения, значению, на котором началась проверка.
*/

func createItemList(prevItem *listNode, Val int) (nextItem *listNode) {
	nextItem = &listNode{Val: Val}
	prevItem.Next = nextItem
	return nextItem
}

func createList(slice []int) *listNode {
	switch len(slice) {
	case 0:
		return nil
	case 1:
		return &listNode{Val: slice[0]}
	case 2:
		nextItem := &listNode{Val: slice[1]}
		return &listNode{Val: slice[0], Next: nextItem}
	default:
		nextItem := &listNode{Val: slice[1]}
		firstItem := &listNode{Val: slice[0], Next: nextItem}
		for i := 2; i < len(slice); i++ {
			nextItem = createItemList(nextItem, slice[i])
		}
		return firstItem
	}
}

func deleteDuplicates(list *listNode) {
	for list.Next != nil {
		if (list.Val == list.Next.Val) && (list.Next != nil) {
			list.Next = list.Next.Next
		} else {
			list = list.Next
		}
	}
}

func main() {
	slice := []int{1, 2, 2, 2, 2, 2, 3, 4, 5, 5, 5}
	list := createList(slice)
	forList := list
	for i := 0; i < len(slice); i++ {
		fmt.Println(forList.Val)
		forList = forList.Next
	}

	deleteDuplicates(list)
	for list.Next != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
