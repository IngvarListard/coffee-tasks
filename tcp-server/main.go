package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	first = iota
	operator
	last
)

var port = flag.String("port", "8001", "the port to listen to")

func convertArgs(exp []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(exp[first], 64)
	if err != nil {
		return .0, .0, fmt.Errorf("unknown number: %v", exp[first])
	}
	num2, err := strconv.ParseFloat(exp[last], 64)
	if err != nil {
		return .0, .0, fmt.Errorf("unknown number: %v", exp[last])
	}
	return num1, num2, nil
}

func resolveExpression(e string) string {

	exp := strings.Split(e, " ")

	if len(exp) != 3 {
		return "error: wrong expression format"
	}
	num1, num2, err := convertArgs(exp)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	res := .0
	switch exp[operator] {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		return "error: operator is not supported"
	}

	return fmt.Sprintf("Result: %g", res)
}

func handleConnection(conn net.Conn) {
	log.Printf("Connected %s\n", conn.RemoteAddr().String())
	defer conn.Close()

	conn.Write([]byte("Hi, this is a calculator application. Enter an expression separated by spaces. " +
		"At the moment, only simple expressions with one operator are supported.\n"))
	conn.Write([]byte("Expression: "))

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			log.Println("error reading data", err)
			return
		}

		exp := strings.TrimSpace(netData)
		if exp == "exit" {
			return
		}
		conn.Write([]byte(resolveExpression(exp) + "\n"))
		conn.Write([]byte("Expression: "))

	}
}

func main() {
	flag.Parse()

	address := "localhost:" + *port
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal("error starting listener:", err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			log.Println("error closing listener:", err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}
		go handleConnection(conn)
	}
}
