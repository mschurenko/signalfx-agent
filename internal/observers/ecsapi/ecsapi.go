// Package ecsapi is a foo-based observer
package ecsapi

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/signalfx/signalfx-agent/internal/core/config"
	"github.com/signalfx/signalfx-agent/internal/core/services"
	"github.com/signalfx/signalfx-agent/internal/observers"
	log "github.com/sirupsen/logrus"
)

const (
	observerType = "ecs-api"
	dockerLabel  = "fargateExporterPort"
)

var logger = log.WithFields(log.Fields{"observerType": observerType})

// Config for the file observer
type Config struct {
	config.ObserverConfig
	ECSClusterName string `yaml:"ECSClusterName"`
}

// ECSAPI observer plugin
type ECSAPI struct {
	serviceCallbacks *observers.ServiceCallbacks
	serviceDiffer    *observers.ServiceDiffer
	config           *Config
}

func init() {
	observers.Register(observerType, func(cbs *observers.ServiceCallbacks) interface{} {
		return &ECSAPI{
			serviceCallbacks: cbs,
		}
	}, &Config{})
}

// Configure foo
func (e *ECSAPI) Configure(config *Config) error {
	e.config = config

	if e.serviceDiffer != nil {
		e.serviceDiffer.Stop()
	}

	e.serviceDiffer = &observers.ServiceDiffer{
		DiscoveryFn:     e.discover,
		IntervalSeconds: 5,
		Callbacks:       e.serviceCallbacks,
	}
	e.serviceDiffer.Start()

	return nil
}

type taskThing struct {
	IP   string
	port uint16
}

func getTasks(clusterName string) []taskThing {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String("us-west-2")},
	))
	svc := ecs.New(sess)

	var tasks []taskThing

	listTasksInput := &ecs.ListTasksInput{Cluster: aws.String(clusterName)}

	listTasksOutput, err := svc.ListTasks(listTasksInput)
	if err != nil {
		log.Fatalln(err)
	}

	if len(listTasksOutput.TaskArns) < 1 {
		return []taskThing{}
	}

	fmt.Println("debug:", "we have task outputs")

	descTasksInput := &ecs.DescribeTasksInput{
		Cluster: aws.String(clusterName),
		Tasks:   listTasksOutput.TaskArns,
	}

	descTasksOutput, err := svc.DescribeTasks(descTasksInput)
	if err != nil {
		log.Fatalln(err)
	}

	// task := descTasksOutput.Tasks[0]
	for _, task := range descTasksOutput.Tasks {
		if *task.LaunchType != "FARGATE" {
			fmt.Println(fmt.Sprintf("info: %v is not FARGATE", *task.TaskArn))
		}

		var taskIP *string
		var taskDef *string
		var taskPort *string

		taskDef = task.TaskDefinitionArn
		details := task.Attachments[0].Details

		for _, detail := range details {
			if *detail.Name == "privateIPv4Address" {
				taskIP = detail.Value
				break
			}
		}

		descTaskDefInput := &ecs.DescribeTaskDefinitionInput{TaskDefinition: taskDef}
		descTaskDefOutput, err := svc.DescribeTaskDefinition(descTaskDefInput)
		if err != nil {
			log.Fatalln(err)
		}

		for _, c := range descTaskDefOutput.TaskDefinition.ContainerDefinitions {
			for k, v := range c.DockerLabels {
				if k == dockerLabel {
					taskPort = v
				}
			}
		}

		p, err := strconv.Atoi(*taskPort)
		if err != nil {
			log.Fatalln(err)
		}

		if taskIP != nil && taskPort != nil {
			fmt.Println("taskIP:", *taskIP)
			fmt.Println("taskPort:", *taskPort)

			tasks = append(tasks, taskThing{
				IP:   *taskIP,
				port: uint16(p),
			})
		}

	}

	return tasks
}

// Discover services from env var
func (e *ECSAPI) discover() []services.Endpoint {
	clusterName := e.config.ECSClusterName

	var instances []*services.ECSAPI
	var out []services.Endpoint

	tasks := getTasks(clusterName)
	if len(tasks) > 0 {
		for _, task := range tasks {

			fargateTask := &services.ECSAPI{
				EndpointCore: *services.NewEndpointCore(task.IP, task.IP, observerType, map[string]string{}),
				IsFargate:    true,
			}

			fargateTask.Port = uint16(task.port)
			fargateTask.PortType = services.TCP
			fargateTask.Host = task.IP

			instances = append(instances, fargateTask)
		}
	}

	for _, i := range instances {
		out = append(out, i)
	}

	return out
}

// Shutdown the service differ routine
func (e *ECSAPI) Shutdown() {
	if e.serviceDiffer != nil {
		e.serviceDiffer.Stop()
	}
}
