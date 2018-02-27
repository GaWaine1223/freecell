# freecell
Just a blockchain, Nothing in particular.

## 二．General design

	main.go // 入口
	server  // 传输层入口 
		- http      // HTTP server to support remote operation
		- command line // standard inputs to support local operation
	handlers        // 函数入口层
		- transaction
            - activity.go // 运营活动与券
            - bill.go       // 账单价格
            - pay.go    // 支付方式等
        - append
	models
		- block
	protocal
	    - status machine
	common
	    - idl
		- errno
		- const
		- util