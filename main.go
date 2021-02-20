package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	action string
	addCmd = flag.NewFlagSet("add", flag.ExitOnError)
	delCmd = flag.NewFlagSet("del", flag.ExitOnError)
	updateCmd = flag.NewFlagSet("update", flag.ExitOnError)
	listCmd = flag.NewFlagSet("list",flag.ExitOnError)
)

const (
	todoFilename =  ".todoList"
)

const (
	doneMark1 = "\u2610"
	doneMark2 = "\u2611"
	successMark  = "√"
	eventMark = "-"
)

func main() {
	existCurTodo := false
	currentPath := ""
	todoPath,err := os.Getwd()
	if err == nil{
		currentPath = filepath.Join(todoPath,todoFilename)
		_,err = os.Stat(currentPath)
		if err == nil{
			existCurTodo = true
		}
	}

	if !existCurTodo{
		// 寻找根目录下的todo文件
		home := os.Getenv("HOME")
		if home == ""{
			home = os.Getenv("USERPROFILE")
		}
		currentPath = filepath.Join(home, todoFilename)
	}


	if len(os.Args) < 2{
		log.Fatal("参数过短")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.StringVar(&action,"name","add","add something")
		addCmd.Parse(os.Args[2:])
		// 合并在拆分
		totalString := strings.Join(addCmd.Args(),"")
		totalStringList := strings.Split(totalString,"-")

		for _,msg := range totalStringList{
			if msg != ""{
				err := addToFile(msg,currentPath)
				if err != nil{
					log.Fatal(err)
				}
			}

		}
	case "del":
		delCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'del'")
	case "update":
		updateCmd.Parse(os.Args[2:])
		fmt.Println("subcommadn 'update'")
	case "list":
		listCmd.Parse(os.Args[2:])
		listFromFile(currentPath)
	default:
		fmt.Println("")
		os.Exit(1)
	}
}

func listFromFile(path string) error{
	//w, err := os.OpenFile(path,os.O_RDONLY,0600)
	r, err := os.Open(path)
	if err != nil{
		log.Fatal(err)
	}
	defer  r.Close()
	br := bufio.NewReader(r)

	n:=1
	isempty := true
	for {
		b, _, err := br.ReadLine()
		if err != nil{
			if err != io.EOF {
				return err
			}
			if isempty{
				fmt.Println("暂时没有日程")
			}
			break
		}
		isempty = false
		line := string(b)
		if strings.HasPrefix(line, "-") {
			fmt.Printf("%s %03d: %s\n", doneMark2, n, strings.TrimSpace(line[1:]))
		} else {
			fmt.Printf("%s %03d: %s\n", doneMark1, n, strings.TrimSpace(line))
		}
		n++
	}
	return nil
}

func addToFile(msg string,path string)error{
	w, err := os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_RDWR,0666)
	if err != nil{
		log.Fatal(err)
	}
	defer  w.Close()
	_, err = fmt.Fprintln(w, msg)
	fmt.Printf("event add:%s%s\n",doneMark1,msg)
	return err
}