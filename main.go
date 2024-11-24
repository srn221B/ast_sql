package main

import (
	"fmt"
	"strings"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/format"
	"github.com/pingcap/tidb/pkg/parser/model"
)

func main() {
	insertStmt := &ast.InsertStmt{
		Table: &ast.TableRefsClause{
			TableRefs: &ast.Join{
				Left: &ast.TableName{
					Name: model.NewCIStr("user"),
				},
			},
		},
		Columns: []*ast.ColumnName{
			{
				Name: model.NewCIStr("name"),
			},
		},
	}
	var sb strings.Builder
	restoreCtx := format.NewRestoreCtx(format.DefaultRestoreFlags, &sb)

	// InsertStmtをSQLに戻す
	err := insertStmt.Restore(restoreCtx)
	if err != nil {
		fmt.Printf("Error restoring SQL: %v\n", err)
		return
	}

	// 復元されたSQLを出力
	fmt.Println("Restored SQL:", sb.String())
}
