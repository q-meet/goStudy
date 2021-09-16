package main

import (
	"frames/example"
)

/*
(选做)本周作业比较大：
•整合 httprouter，
•整合 validator，和 gin 的 binding 库
•实现兼容 http.HandlerFunc 的 middleware

完成一个属于自己的 restful ⻛格 API 框架
*/

func main() {
	//httprouter 路由 example
	example.Router()

	//整合 validator example  middleware只是包了一层
	example.MyValidate()

	//整合 sqlx example
	example.MysqlSql()
	example.PostgresSql()
}
