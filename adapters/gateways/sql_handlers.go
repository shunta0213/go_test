/*
DIP (Dependency Inversion Principal)
Avoid using methods implemented in infra layer in gateways (adaptors) layer.
*/
/*
In the future, We need to fix user_gateways following to DIP.
*/

package gateways

// type SqlHandler interface {
// 	Execute(string, ...interface{}) (Result, error)
// 	Query(string, ...interface{}) (Row, error)
// }

// type Result interface {
// 	LastInsertId() (int64, error)
// 	RowsAffected() (int64, error)
// }

// type Row interface {
// 	Scan(...interface{}) error
// 	Next() bool
// 	Close() error
// }
