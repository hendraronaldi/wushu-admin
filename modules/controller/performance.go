package controller

import (
	"encoding/json"
	"fmt"
	"strings"
	"work/wushu-backend/modules/connections"
	"work/wushu-backend/modules/model"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func PostPerformance(c *gin.Context) {
	var data model.Performance
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid login request",
		})
	} else {
		data.Email = strings.ToLower(data.Email)
		if _, err = FindUser(data.Email); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.PostgresConnection()
			if conn != nil {
				performance := structs.Map(data)
				details := structs.New(data)
				performanceCategories := structs.Names(data)

				// Insert Performance Details
				for _, category := range performanceCategories {
					if category != "Email" && category != "Date" && !strings.HasSuffix(category, "_id") {
						var columnNames []string
						var detailValues []interface{}

						tableName := strings.ToLower(category)

						for _, field := range details.Field(category).Fields() {
							columnNames = append(columnNames, strings.ToLower(field.Name()))
							detailValues = append(detailValues, field.Value())
						}

						addedID := connections.InsertPostgresData(conn, tableName, columnNames, detailValues)
						performance[strings.Title(category)+"_id"] = addedID
					}
				}

				// Inser Performance Overall
				var columns []string
				var values []interface{}

				for key, val := range performance {
					if key == "Email" || key == "Date" || strings.HasSuffix(key, "_id") {
						columns = append(columns, strings.ToLower(key))
						values = append(values, val)
					}
				}

				fmt.Println(columns)
				fmt.Println(values)

				performanceID := connections.InsertPostgresData(conn, "performance", columns, values)

				if performanceID == 0 {
					c.JSON(400, gin.H{
						"response": "user not exist",
					})
				} else {
					c.JSON(200, gin.H{
						"response": "success",
					})
				}
			} else {
				c.JSON(400, gin.H{
					"response": "connection fail",
				})
			}
		}
	}
}
