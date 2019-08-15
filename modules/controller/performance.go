package controller

import (
	"encoding/json"
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

				performanceID := connections.InsertPostgresData(conn, "performance", columns, values)

				if performanceID == 0 {
					c.JSON(400, gin.H{
						"response": "fail to insert data",
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

func GetPerformance(c *gin.Context) {
	var performance model.Performance
	var performances []interface{}
	var unusedID string

	conn := connections.PostgresConnection()
	query := `SELECT * FROM performance, flexibility, power
	WHERE performance.flexibility_id = flexibility.flexibility_id
	AND performance.power_id = power.power_id
	AND performance.performance_id > 0
	AND performance.flexibility_id > 0
	AND performance.power_id > 0`

	rows, err := conn.Query(query)
	if err != nil {
		// handle this error better than this
		c.JSON(400, gin.H{
			"response": "query error",
		})
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&performance.Date, &performance.Flexibility_id, &performance.Email, &unusedID, &performance.Power_id,
			&unusedID, &performance.Flexibility.Shoulder, &performance.Flexibility.Wrist,
			&performance.Flexibility.Waist, &performance.Flexibility.Leg,
			&unusedID, &performance.Power.Jump, &performance.Power.Kick,
			&performance.Power.Strike, &performance.Power.HandSwing, &performance.Power.Spin, &performance.Power.LegSwing)
		if err != nil {
			// handle this error
			c.JSON(400, gin.H{
				"response": "bind data error",
			})
		}
		performances = append(performances, performance)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		c.JSON(400, gin.H{
			"response": "iteration error",
		})
	} else {
		c.JSON(200, performances)
	}
}

func GetUserPerformance(c *gin.Context) {
	var performance model.Performance
	var performances []interface{}
	var unusedID string
	email := c.Param("email")

	conn := connections.PostgresConnection()

	query := `SELECT * FROM performance, flexibility, power
	WHERE performance.flexibility_id = flexibility.flexibility_id
	AND performance.email LIKE '` + email + `'
	AND date_part('year', performance.date) = date_part('year', current_date)
	AND performance.power_id = power.power_id
	AND performance.performance_id > 0
	AND performance.flexibility_id > 0
	AND performance.power_id > 0`

	rows, err := conn.Query(query)
	if err != nil {
		// handle this error better than this
		c.JSON(400, gin.H{
			"response": "query error",
		})
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&performance.Date, &performance.Flexibility_id, &performance.Email, &unusedID, &performance.Power_id,
			&unusedID, &performance.Flexibility.Shoulder, &performance.Flexibility.Wrist,
			&performance.Flexibility.Waist, &performance.Flexibility.Leg,
			&unusedID, &performance.Power.Jump, &performance.Power.Kick,
			&performance.Power.Strike, &performance.Power.HandSwing, &performance.Power.Spin, &performance.Power.LegSwing)
		if err != nil {
			// handle this error
			c.JSON(400, gin.H{
				"response": "bind data error",
			})
		}
		performances = append(performances, performance)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil || performances == nil {
		c.JSON(400, gin.H{
			"response": "iteration error",
		})
	} else {
		c.JSON(200, performances)
	}
}
