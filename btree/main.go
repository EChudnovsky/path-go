/* Бинарное дерево поиска (BST) в Go
BST — это бинарное дерево, где:
• Значение в левой ветви меньше значения в узле.
• Значение в правой ветви больше значения в узле.

Такое свойство обеспечивает эффективный поиск, вставку и удаление — в среднем за O(log n).

Как это работает?
1. Вставка:
• Если значение меньше текущего узла, идём в левое поддерево.
• Если больше — в правое.
• Если узел пуст, создаём новый.

2. Поиск:
• Сравниваем искомое значение с текущим.
• Меньше — идём влево, больше — вправо, равное — нашли.
3. Обход:

• In-order обход (LNR) обходит узлы в отсортированном порядке.
*/

package main

// Node — узел бинарного дерева
type Node struct {
	Value int
	Left *Node
	Rigth *Node
}

NewNode() *Node {
	return &Node{
		NewProduct(name, "Watersports", price),
		capacity,
		motorized,
	}
}


// Insert — вставка нового значения в дерево
func (n *Node) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node {
				Value: value,
			}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &Node {
				Value: value,
			}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search — поиск значения в дереве
func (n *Node) Search(value int) *Node {
	if n == nil {
		return nil
	}
	if n.Value == value {
		return n
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// InOrderTraversal — обход дерева в порядке (LNR)
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Print(n.Value, " ")
	n.Right.InOrderTraversal()
}

func main() {
	bt := Node{}
	bt.Insert(1)

}