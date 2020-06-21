package main

import (
	"fmt"
	"github.com/IngvarListard/coffee-tasks/binary-protocols/prot1"
	"log"
	"time"
)

func main() {
	table := prot1.Table{
		Headers: []string{"user_id", "salary", "parent_id", "full_name", "time"},
		Types:   []uint{prot1.Int32Code, prot1.Float64Code, prot1.BoolCode, prot1.StringCode, prot1.TimeCode},
		Data: []interface{}{
			[]interface{}{int32(1), 33.82, false, "Иванов Иван Иваныч", time.Now()},
			[]interface{}{int32(2), 18.44, true, "Козлов Петр Петрович", time.Now()},
			[]interface{}{int32(3), 14.88, false, "Калов Кал Калыч", time.Now()},
			[]interface{}{int32(4), 9.08, true, "Григорьев Ибрагим Коловратович", time.Now()},
			[]interface{}{int32(5), 1.08, false, "Залетов Петр Иосифович", time.Now()},
			[]interface{}{int32(6), 88.55, true, "Глинин Камень Горыныч", time.Now()},
		},
	}
	data, err := table.MarshalBinary()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("BYTES", len(data))

	deserializedTable := new(prot1.Table)
	err = deserializedTable.UnmarshalBinary(data)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range deserializedTable.Data {
		fmt.Println(v)
	}
}
