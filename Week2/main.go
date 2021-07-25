// 当dao层中遇到一个sql.ErrNoRows时，不应该Wrap这个error，应该直接往上抛
// 因为dao层属于基础库，很多上层业务调用，如果dao层wrap这个错误，业务也wrap错误，导致日志记录两遍堆栈信息
// 在第三节课，Handing Error提到，只有Application层需要Wrap error
package main

import (
	"database/sql"
)

var (
	db *sql.DB
)

func QueryUsernameById(id int) (error, string) {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id=?", id).Scan(&username)
	if err != nil {
		return err, username
	}
	return nil, username
}
