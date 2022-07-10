package main

import (
	"flag"
	"fmt"
	"net/http"
	"trino.com/trino-connectors/DAO"
	"trino.com/trino-connectors/data"
	"trino.com/trino-connectors/service"
	"trino.com/trino-connectors/util/log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// InitConfig()
	// err := DAO.InitMysqlConfig()
	// if err != nil {
	// 	return
	// }
	//DAO.InitRedisConfig()

	DAO.InitialPostgres()
	testPort := flag.String("port", "8081", "The port to use")
	flag.Parse()

	router := gin.Default()
	router.Use(log.LoggerToFile())

	logger := log.Logger()

	// GET /v1/connector/list_connections
	router.POST(service.GetNoDBDataPath, func(c *gin.Context) {
		req := data.GetDataReq{}
		err := c.BindJSON(&req)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		d, err := service.GetData(req.TableName, req.FieldName)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		resp := data.GetData{
			Ret:  0,
			Msg:  "Get Data successfully.",
			Data: d,
		}

		c.JSON(http.StatusOK, resp)
		logger.WithField("path", service.ListConnectionsPath).Info("request_processed_successfully")
	})

	fmt.Println("program start")
	err := router.Run(fmt.Sprintf(service.TestPort, *testPort))
	if err != nil {
		return
	}
}
