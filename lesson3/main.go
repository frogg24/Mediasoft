package main

import (
	"fmt"
	"mediasoft/lesson3/binaryTree"
	"mediasoft/lesson3/linkedList"
	"mediasoft/lesson3/stack"
)

func main() {
	showWork()
}

func showWork() {
	fmt.Println("Демонстрация работы")

	//Работа со стеком
	fmt.Println("Стек:")
	stackInt := stack.NewStack[int](10)
	stackInt.Push(1)
	stackInt.Push(3)
	stackInt.Push(5)
	stackInt.Push(10)
	fmt.Println("Врехнее значение:", stackInt.Peek())
	if value, ok := stackInt.Pop(); ok {
		fmt.Println("Верхнее значение (удалено):", value)
	} else {
		fmt.Println("Стек пуст")
	}
	if value, ok := stackInt.Pop(); ok {
		fmt.Println("Верхнее значение (удалено):", value)
	} else {
		fmt.Println("Стек пуст")
	}
	fmt.Println("Врехнее значение:", stackInt.Peek())
	stackString := stack.NewStack[string](10)
	stackString.Push("1st")
	stackString.Push("2nd")
	fmt.Println("Врехнее значение:", stackString.Peek())
	if value, ok := stackString.Pop(); ok {
		fmt.Println("Верхнее значение (удалено):", value)
	} else {
		fmt.Println("Стек пуст")
	}
	if value, ok := stackString.Pop(); ok {
		fmt.Println("Верхнее значение (удалено):", value)
	} else {
		fmt.Println("Стек пуст")
	}
	if value, ok := stackString.Pop(); ok {
		fmt.Println("Верхнее значение (удалено):", value)
	} else {
		fmt.Println("Стек пуст")
	}

	//Связный список
	fmt.Println("Связный список:")
	listInt := linkedList.NewLinkedList[int]()
	listInt.Add(10)
	listInt.Add(20)
	listInt.Add(30)
	listInt.Add(40)
	fmt.Println("Получено значение: ", listInt.Get(2))
	listInt.Remove(1)
	fmt.Println("Связный список: ", listInt.Values())

	//Бинарное дерево
	fmt.Println("Бинарное дерево:")
	treeInt := binaryTree.NewBinaryTree[int]()
	treeInt.Add(50)
	treeInt.Add(40)
	treeInt.Add(70)
	treeInt.Add(60)
	treeInt.Add(80)
	treeInt.Add(65)
	fmt.Println("Бинарное дерево: ", treeInt.Values())
	treeInt.Remove(70)
	fmt.Println("Бинарное дерево: ", treeInt.Values())
	treeFloat := binaryTree.NewBinaryTree[float32]()
	treeFloat.Add(0.50)
	treeFloat.Add(0.40)
	treeFloat.Add(0.70)
	treeFloat.Add(0.60)
	treeFloat.Add(0.80)
	treeFloat.Add(0.65)
	fmt.Println("Бинарное дерево: ", treeFloat.Values())
	treeFloat.Remove(0.70)
	fmt.Println("Бинарное дерево: ", treeFloat.Values())
	treeFloat.Remove(0.50)
	treeFloat.Remove(0.40)
	treeFloat.Remove(0.65)
	treeFloat.Remove(0.80)
	fmt.Println("Бинарное дерево: ", treeFloat.Values())
	treeFloat.Remove(0.60)
	fmt.Println("Бинарное дерево: ", treeFloat.Values())
}
