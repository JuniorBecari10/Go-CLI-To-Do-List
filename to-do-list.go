package main

import (
    "fmt"
    "bufio"
    "os"
    "os/exec"
    "runtime"
    "strconv"
    "strings"
)

type Task struct {
    name string
    completed bool
}

var clear map[string]func() //create a map for storing clear funcs

// --------

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var list []Task

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func printList(list []Task) {
    if len(list) == 0 {
        fmt.Println("List is empty. Try adding one task.")
    } else {
        fmt.Println("Name | Completed")
        
        for i := 0; i < len(list); i++ {
            fmt.Println(list[i].name, "|", boolToYes(list[i].completed))
        }
    }
}

func boolToYes(b bool) string {
    if b {
        return "Yes"
    } else {
        return "No"
    }
}

func removeItem(list []Task, index int64) []Task {
    ret := make([]Task, 0)
    ret = append(ret, list[:index]...)
    return append(ret, list[index + 1:]...)
}

// --------

func addTask(name string) {
    output := make([]Task, len(list) + 1)
    output = append(list, Task{name, false})
    
    list = output
}

func printOptions() {
    if len(list) == 0 {
        fmt.Println("1. Add Task")
        fmt.Println("2. Exit")
    } else {
        fmt.Println("1. Add Task")
        fmt.Println("2. Remove Task")
        fmt.Println("3. Set Completed/Uncompleted")
        fmt.Println("4. Rename Task")
        if len(list) >= 2 {
            fmt.Println("5. Swap Tasks")
        } else {
            fmt.Println("5. Swap Tasks | DEACTIVATED - List is too short")
        }
        
        fmt.Println("6. Clear List")
        fmt.Println("7. Exit")
    }
}

func swapIndex(list []Task, a int, b int) []Task {
    indexA := list[a]
    
    list[a] = list[b]
    list[b] = indexA
    
    return list
}

func processInput(input string) {
    if len(list) == 0 {
        switch input {
            case "1":
                fmt.Printf("Enter name: ")
                scanner.Scan()
                
                name := scanner.Text()
                
                addTask(name)
            case "2":
                os.Exit(0)
        }
    } else {
        switch input {
            case "1":
                fmt.Printf("Enter name: ")
                scanner.Scan()
                
                name := scanner.Text()
                
                addTask(name)
            case "2":
                fmt.Printf("Enter index: ")
                scanner.Scan()
                
                index, _ := strconv.ParseInt(scanner.Text(), 10, 0)
                
                if int(index) >= len(list) {
                    return
                }
                
                list = removeItem(list, index)
            case "3":
                fmt.Printf("Enter index: ")
                scanner.Scan()
                
                index, _ := strconv.ParseInt(scanner.Text(), 10, 0)
                
                if int(index) >= len(list) {
                    return
                }
                
                list[index].completed = !list[index].completed
            case "4":
                fmt.Printf("Enter index: ")
                scanner.Scan()
                
                index, _ := strconv.ParseInt(scanner.Text(), 10, 0)
                
                if int(index) >= len(list) {
                    return
                }
                
                fmt.Printf("Enter new name: ")
                scanner.Scan()
                
                name := scanner.Text()
                
                list[index].name = name
            case "5":
                if len(list) < 2 {
                    return
                }
                
                fmt.Printf("Enter index 1: ")
                scanner.Scan()
                
                indexA, _ := strconv.ParseInt(scanner.Text(), 10, 0)
                
                if int(indexA) >= len(list) {
                    return
                }
                
                fmt.Printf("Enter index 2: ")
                scanner.Scan()
                
                indexB, _ := strconv.ParseInt(scanner.Text(), 10, 0)
                
                if int(indexB) >= len(list) {
                    return
                }
                
                list = swapIndex(list, int(indexA), int(indexB))
           case "6":
               fmt.Printf("Are you sure? (y/n) ")
               scanner.Scan()
               
               sure := strings.ToLower(scanner.Text())
               
               if sure == "y" || sure == "yes" {
                   list = make([]Task, 0)
               }
           case "7":
                os.Exit(0)
        }
    }
}

func main() {
    for true {
        CallClear()
        fmt.Printf("To-Do-List\n\n")
        
        printList(list)
        fmt.Println()
        
        printOptions()
        fmt.Println()
        
        fmt.Printf("Enter input: ")
        scanner.Scan()
        
        input := scanner.Text()
        
        processInput(input)
    }
}
