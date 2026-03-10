package main

import (
	"context"
	"fmt"
	simpleconnection "study/feature_postgres/simple_connection"
	simplesql "study/feature_postgres/simple_sql"
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
	/*if err := simplesql.InsertRow(ctx,
		conn,
		"Обед2",
		"Покушать надо",
		false,
		time.Now(),
	); err != nil {
		panic(err)
	}
	*/

	if err := simplesql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}

	/*if err := simplesql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}
	*/
	fmt.Println("succeed")
}
