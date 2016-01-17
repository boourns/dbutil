// Generated by ego.
// DO NOT EDIT

package dbutil
import (
"fmt"
"html"
"io"
"reflect"
)
var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
//line create_table.sql.ego:1
 func createTable(w io.Writer, t reflect.Type) error  {
//line create_table.sql.ego:2
_, _ = io.WriteString(w, "\n\n")
//line create_table.sql.ego:4
_, _ = io.WriteString(w, "\n\nCREATE TABLE ")
//line create_table.sql.ego:5
_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v",  t.Name() )))
//line create_table.sql.ego:5
_, _ = io.WriteString(w, " (\n  ")
//line create_table.sql.ego:6
 for i := 0; i < t.NumField(); i++ { 
//line create_table.sql.ego:7
_, _ = io.WriteString(w, "\n    ")
//line create_table.sql.ego:7
 field := t.Field(i) 
//line create_table.sql.ego:8
_, _ = io.WriteString(w, "\n    ")
//line create_table.sql.ego:8
_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v",  field.Name )))
//line create_table.sql.ego:8
_, _ = io.WriteString(w, " ")
//line create_table.sql.ego:8
_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v",  sqlType(field) )))
//line create_table.sql.ego:8
 if i < t.NumField()-1 { 
//line create_table.sql.ego:8
_, _ = io.WriteString(w, ",")
//line create_table.sql.ego:8
 }
//line create_table.sql.ego:9
_, _ = io.WriteString(w, "\n  ")
//line create_table.sql.ego:9
 } 
//line create_table.sql.ego:10
_, _ = io.WriteString(w, "\n);\n\n")
return nil
}