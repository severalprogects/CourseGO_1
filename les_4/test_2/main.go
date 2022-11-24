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
		// return createItemList(&listNode{Val: slice[0]}, slice[1]) //Возвращает последний элемент, а не первый
	default:
		/*
			nextItem := &listNode{Val: slice[1]}
			firstItem := &listNode{Val: slice[0], Next: nextItem}
			prevItem := nextItem

			for i := 2; i < len(slice); i++ {
				nextItem = &listNode{Val: slice[i]}
				prevItem.Next = nextItem
			}
			return firstItem
		*/

		nextItem := &listNode{Val: slice[1]}
		firstItem := &listNode{Val: slice[0], Next: nextItem}
		for i := 1; i < len(slice); i++ {
			nextItem = createItemList(nextItem, slice[i])
		}
		return firstItem
	}
}

func main() {
	slice := []int{1, 2, 2, 2, 2, 2, 3, 4, 5, 5, 5}
	list := createList(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println(list.Val)
		list = list.Next
	}
}
