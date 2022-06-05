package controller

//func Feed(c *gin.Context) {
//	// latestTime
//	var latestTime time.Time
//
//	strTime := c.Query("latest_time")
//	if strTime == "" {
//		latestTime = time.Now()
//
//	} else {
//		last_time_unix, err := strconv.ParseInt(strTime, 10, 64)
//		if err != nil {
//			c.JSON(http.StatusOK, handlers.Response{
//				StatusCode: 500,
//				StatusMsg:  "last_time时间戳错误",
//			})
//			return
//		}
//		latestTime = time.Unix(last_time_unix/1000, 0)
//	}
//
//	// token
//	claims := jwt.ExtractClaims(c)
//	userID := int64(claims[constants.IdentityKey].(float64))
//
//	// todo 封装数据返回
//
//}
