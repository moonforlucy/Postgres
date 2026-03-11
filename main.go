package main

import (
	"context"
	"fmt"
	simpleconnection "study/feature_postgres/simple_connection"
	simplesql "study/feature_postgres/simple_sql"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simplesql.CrateTable(ctx, conn); err != nil {
		panic(err)
	}
	/*
		if err := simplesql.InsertRow(ctx,
			conn,
			"Прогулка с собакой",
			"Прогуляться с собакой после работы",
			false,
			time.Now(),
		); err != nil {
			panic(err)
		}
	*/
	/*if err := simplesql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}
	*/
	/*if err := simplesql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}
	*/
	tasks, err := simplesql.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}
	pp.Println(tasks)
	fmt.Println("succeed")
}
