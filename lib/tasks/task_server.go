package tasks

import (
	"encoding/json"
	"log"

	"github.com/jivison/gowon-indexer/lib/db"

	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/tasks"

	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
)

// TaskServer is an interface to manage tasks
var TaskServer *GowonTaskServer

type GowonTaskServer struct {
	Server  *machinery.Server
	Workers []*machinery.Worker
}

// SendIndexUserTask starts an index task
func (gts GowonTaskServer) SendIndexUserTask(user *db.User, token string) {
	bytesArray, _ := json.Marshal(user)

	signature := &tasks.Signature{
		Name: "index_user",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: string(bytesArray),
			},
			{
				Type:  "string",
				Value: token,
			},
		},
	}

	gts.Server.SendTask(signature)
}

// SendUpdateUserTask starts an update task
func (gts GowonTaskServer) SendUpdateUserTask(user *db.User, token string) {
	bytesArray, _ := json.Marshal(user)

	signature := &tasks.Signature{
		Name: "update_user",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: string(bytesArray),
			},
			{
				Type:  "string",
				Value: token,
			},
		},
	}

	gts.Server.SendTask(signature)
}

// LaunchWorkers launches all the workers associated with a task server
func (gts GowonTaskServer) LaunchWorkers() {
	for _, worker := range gts.Workers {
		go launchWorker(worker)
	}
}

// NewTaskServer creates a new task
func NewTaskServer() *GowonTaskServer {
	if TaskServer != nil {
		return TaskServer
	}

	server := createServer()
	workers := []*machinery.Worker{
		server.NewWorker("lastfm_worker", 5),
	}

	TaskServer = &GowonTaskServer{
		Server:  server,
		Workers: workers,
	}

	return TaskServer
}

func createServer() *machinery.Server {
	cnf := &config.Config{
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	broker := redisbroker.NewGR(cnf, []string{"localhost:6379"}, 0)
	backend := redisbackend.NewGR(cnf, []string{"localhost:6379"}, 0)
	lock := eagerlock.New()

	server := machinery.NewServer(cnf, broker, backend, lock)

	server.RegisterTask("index_user", IndexUserTask)
	server.RegisterTask("update_user", UpdateUserTask)

	return server
}

func launchWorker(worker *machinery.Worker) {
	err := worker.Launch()

	if err != nil {
		log.Fatal("Error creating machinery worker: ", err)
	}
}
