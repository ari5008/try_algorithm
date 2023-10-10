package main

import (
    "backend/try/stack_que"
    "fmt"
)

func main() {
    stack := &try.Stack{}
    fmt.Println(stack.GetStack())
    
    stack.Push(1)
    fmt.Println(stack.GetStack())
    
    stack.Push(2)
    fmt.Println(stack.GetStack())
    
    fmt.Println(stack.Pop())
    fmt.Println(stack.GetStack())
    
    fmt.Println(stack.Pop())
    fmt.Println(stack.GetStack())
}