package main

//func main() {
//	r := gin.Default()
//	// json
//	r.GET("/someJSON", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
//	})
//	// struct
//	r.GET("/someStruct", func(c *gin.Context) {
//		var msg struct {
//			Name    string
//			Message string
//			Number  int
//		}
//		msg.Name = "root"
//		msg.Message = "message"
//		msg.Number = 123
//		c.JSON(200, msg)
//	})
//	// xml
//	r.GET("/someXML", func(c *gin.Context) {
//		c.XML(200, gin.H{"message": "abc"})
//	})
//	// YAML
//	r.GET("/someYML", func(c *gin.Context) {
//		c.YAML(200, gin.H{"name": "刘君"})
//	})
//	r.Run(":8080")
//}
