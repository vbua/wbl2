package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).
*/

func showProcesses() {
	processList, err := ps.Processes()
	if err != nil {
		fmt.Printf("ps: %v\n", err)
		return
	}
	for x := range processList {
		var process ps.Process
		process = processList[x]
		fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}
}

func killProcess(command string) {
	pid, err := strconv.Atoi(command)
	if err != nil {
		fmt.Printf("kill: %v\n", err)
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("kill: %v\n", err)
		return
	}
	err = process.Kill()
	if err != nil {
		fmt.Printf("kill: %v\n", err)
		return
	}
}

func doFork() {
	pid, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		fmt.Printf("fork: %v\n", err)
		return
	}
	if pid > 0 {
		fmt.Printf("pid: %v\n", pid)
	}
}

func doExec(path string, args []string) {
	binary, err := exec.LookPath(path)
	if err != nil {
		fmt.Printf("exec: %v\n", err)
	}
	env := os.Environ()
	err = syscall.Exec(binary, args, env)
	if err != nil {
		fmt.Printf("exec: %v\n", err)
	}
}

func main() {
	fmt.Println("Welcome to WB shell")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		switch commands[0] {
		case "cd":
			if len(commands) > 1 {
				err := os.Chdir(commands[1])
				if err != nil {
					fmt.Printf("cd: %v\n", err)
				}
			}
		case "pwd":
			path, err := os.Getwd()
			if err != nil {
				fmt.Printf("pwd: %v\n", err)
			}
			fmt.Println(path)
		case "echo":
			if len(commands) > 1 {
				fmt.Println(strings.Join(commands[1:], " "))
			}
		case "kill":
			if len(commands) > 1 {
				killProcess(commands[1])
			}
		case "ps":
			showProcesses()
		case "fork":
			doFork()
		case "exec":
			if len(commands) > 1 {
				doExec(commands[1], commands[1:])
			}
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
}
