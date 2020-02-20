package main 
 
import ( 
   "fmt" 
//   "bufio" 
//   "os" 
//   "math" 
) 

type node struct{
    val int
    next *node
}

func NewNode(value int)*node{
    return &node{
        val: value,
        next: nil,
    }
}

type linkedList struct{
    head *node
}

func NewLinkedList()*linkedList{
    return &linkedList{
        head: nil,
    }
}

func (sll *linkedList)insertAtBeg(value int)*linkedList{
    newNode := NewNode(value)
    if sll.head == nil{
        sll.head = newNode
        return sll
    }
    newNode.next = sll.head
    sll.head = newNode
    return sll
}

func (sll *linkedList)insertAtEnd(value int)*linkedList{
	newNode := NewNode(value)
	if sll.head == nil{
        sll.head = newNode
        return sll
    }
	temp := sll.head
	for temp.next != nil{
		temp = temp.next
	}
	temp.next = newNode
	return sll
}



func (sll *linkedList)traverse(){
    if sll.head == nil{
        return
    }
    temp := sll.head
    for temp != nil{
        fmt.Println(temp.val)
        temp = temp.next
    }
}

// func (sll *linkedList)readFromFront()*node{
// 	if sll.head == nil{
// 		return nil
// 	}
// 	temp := sll.head
// 	for temp != nil{

// 	}
// }
func (sll *linkedList)nthElementInSLL(n int)*node{
	if sll.head == nil{
		return nil
	}
	temp := sll.head
	for i:=1;i<=n && temp != nil;i++{
		temp = temp.next
	}
	return temp
}

func (sll *linkedList)customSLL()*linkedList{
	if sll.head == nil{
		return nil
	}
	newSLL := NewLinkedList()

	newtemp := sll.head
	for newtemp != nil{
		newSLL.insertAtEnd(newtemp.val)
		newtemp = newtemp.next
	}
	
	
	count := 0
	//front := sll.head
	rear := newSLL.head
	for rear.next != nil{
		rear = rear.next
		count++
	}
	custom := NewLinkedList()
	frontCount := 0
	rearCount := count
	for i:=1;i<= count+1;i++{
		if i % 2 != 0{
			custom.insertAtEnd(sll.nthElementInSLL(frontCount).val)
			frontCount++
		}else{
			custom.insertAtEnd(sll.nthElementInSLL(rearCount).val)
			rearCount--
		}
	}
	return custom
}


 
func main() { 
     // INPUT: var n int, arr := make([]int, n), op := 0, l := len(arr), text := "" 
     // READ: io := bufio.NewReader(os.Stdin), fmt.Fscan(io, &n), fmt.Fscan(io, &arr[i]), 
     //       fmt.Scan(&arr[i]), fmt.Scan(&n), fmt.Scanln(text) 
     // DEBUG: fmt.Printf(n), log.Println(n) 
     // OUTPUT: fmt.Println(n), fmt.Println(text), fmt.Println(iop - qop), fmt.Println(operations(s)) 
     // Write the code to solve the problem below, 
     // format the "result" as specified in the problem statement 
     // and finally, write the result to stdout 
	 // IMPORTANT: Remove all debug statements for final submission 
	 
	 //count := 8 // update this value dynamically
     sll := NewLinkedList()
     sll.insertAtEnd(1)
     sll.insertAtEnd(2)
	 sll.insertAtEnd(3)
	 sll.insertAtEnd(4)
	 sll.insertAtEnd(5)
	 sll.insertAtEnd(6)
	 sll.insertAtEnd(7)
	 sll.insertAtEnd(8)
	 
	 fmt.Println("Input SLL: ")
	 sll.traverse()
	 fmt.Println("Customized SLL: ")
	 c := sll.customSLL()
	 c.traverse()
     
}