package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

func main() {
	fmt.Println("\n\n***** Welcome to Linked List Data Structures ****")
	fmt.Println()
	var list linkedList = linkedList{}
	for i := 0; i < 3; i++ {
		list.add_element(i)
		list.add_element(99, i)
		list.add_element(69, list.length)
	}
	fmt.Println("\nAdding Elements:")
	list.print_linked_list()
	fmt.Println("\nSearching Elements:")
	list.search_element(8)
	list.search_element(69)
	fmt.Println("\nGetting Elements:")
	list.get_element(0)
	list.get_element(8)
	list.get_element(100)
	fmt.Println("\nDeleting Elements:")
	list.delete_element(2)
	list.delete_element(99)
	list.delete_element(69)
	list.print_linked_list()
	list.delete_element_at_position(0)
	list.delete_element_at_position(1)
	list.delete_element_at_position(list.length)
	list.delete_element_at_position(69)
	list.print_linked_list()
	fmt.Println("\nReversing Elements:")
	list.reverse_linkedlist()
	list.print_linked_list()
}

func (list *linkedList) add_element(value int, parameters ...int) {
	var element *node = new(node)
	element.value = value
	if len(parameters) != 0 && parameters[0] != 0 {
		var position int = parameters[0]
		var currentPosition int = 0
		for iterator := list.head; iterator != nil; iterator = iterator.next {
			if currentPosition+1 == position {
				element.next = iterator.next
				iterator.next = element
			}
			currentPosition += 1
		}
		if currentPosition < position {
			fmt.Println("Cannot Put element at", position, "due to less list size")
		}
	} else {
		if list.head == nil {
			list.head = element
		} else {
			element.next = list.head
			list.head = element
		}
	}
	list.length += 1
}

func (list linkedList) print_linked_list() {
	if list.head != nil {
		fmt.Print("Printing linked list: ")
		var iterator *node = list.head
		for ; iterator.next != nil; iterator = iterator.next {
			fmt.Print(iterator.value, " --> ")
		}
		fmt.Println(iterator.value)
	} else {
		fmt.Println("List is empty")
	}
}

func (list linkedList) search_element(value int) {
	var found bool
	for iterator := list.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			found = true
			break
		}
	}
	if found {
		fmt.Println(value, "is present in list")
	} else {
		fmt.Println(value, "is not present in list")
	}
}

func (list linkedList) get_element(position int) {
	if list.length < position {
		low_length()
	} else {
		var currentPosition int = 0
		for iterator := list.head; iterator != nil; iterator = iterator.next {
			if currentPosition == position {
				fmt.Println("Value of element at", position, "of the list is", iterator.value)
				break
			}
			currentPosition += 1
		}
	}
}

func (list *linkedList) delete_element(value int) {
	if list.head != nil {
		var previous *node
		for iterator := list.head; iterator != nil; iterator = iterator.next {
			if iterator.value == value {
				if iterator == list.head {
					list.head = iterator.next
				} else {
					previous.next = iterator.next
				}
				break
			}
			previous = iterator
		}
	} else {
		fmt.Println("List is empty")
	}

}

func (list *linkedList) delete_element_at_position(position int) {
	if position > list.length {
		low_length()
	} else {
		var currentPosition int = 0
		var previous *node
		for iterator := list.head; iterator != nil; iterator = iterator.next {
			if currentPosition == position {
				if currentPosition == 0 {
					list.head = iterator.next
				} else {
					previous.next = iterator.next
				}
				break
			}
			previous = iterator
			currentPosition += 1
		}
	}
}

func low_length() {
	fmt.Println("Given list is smaller the size than given index position.")
}

func (list *linkedList) reverse_linkedlist() {
	var current, previous, next *node
	current = list.head
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	list.head = previous
}
