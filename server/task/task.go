package task

import (
	"mt-scale/models"

	"github.com/robfig/cron"
)

var c = cron.New()

func init() {
	// Check mongo database connect
	c.AddFunc("0/60 * * * * ?", models.MongoDbCheck)
}

// Start Start task list
func Start() {
	c.Start()
}
