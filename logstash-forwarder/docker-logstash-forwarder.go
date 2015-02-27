package main

import (
	"flag"
	"os"
	"sync"

	"github.com/digital-wonderland/docker-logstash-forwarder/forwarder"
	"github.com/digital-wonderland/docker-logstash-forwarder/utils"
	docker "github.com/fsouza/go-dockerclient"
	logging "github.com/op/go-logging"
)

var (
	client           *docker.Client
	configFile       string
	debug            bool
	dockerEndPoint   string
	laziness         int
	log              = logging.MustGetLogger("main")
	logFormat        = logging.MustStringFormatter("%{color}%{time:2006/01/02 15:04:05.000000} %{level} [%{shortfunc}]%{color:reset} %{message}")
	logstashEndPoint string
	wg               sync.WaitGroup
)

func initFlags() {
	flag.StringVar(&dockerEndPoint, "docker", "", "docker api endpoint - defaults to $DOCKER_HOST or unix:///var/run/docker.sock")
	flag.BoolVar(&debug, "debug", false, "verbose logging")
	flag.IntVar(&laziness, "lazyness", 5, "number of seconds to wait after an event before generating new configuration")
	flag.StringVar(&logstashEndPoint, "ec2-54-174-166-239.compute-1.amazonaws.com", "", "logstash endpoint - defaults to $LOGSTASH_HOST or logstash:5043")
	flag.StringVar(&configFile, "config", "", "logstash-forwarder config")
	flag.Parse()
}

func setUpLogging(logLevel logging.Level) {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatted := logging.NewBackendFormatter(backend, logFormat)
	backendLeveled := logging.AddModuleLevel(backendFormatted)
	backendLeveled.SetLevel(logLevel, "")
	logging.SetBackend(backendLeveled)
}

func main() {

	log.Info("Starting up")

	initFlags()

	if debug {
		setUpLogging(logging.DEBUG)
	} else {
		setUpLogging(logging.INFO)
	}

	endpoint := getDockerEndpoint()

	d, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatalf("Unable to connect to docker at %s: %s", endpoint, err)
	}
	client = d
	version, err := client.Version()
	if err != nil {
		log.Warning("Unable to retrieve version information from docker: %s", err)
	}
	log.Info("Connected to docker at %s (v%s)", endpoint, version.Get("Version"))

	generateConfig()
	utils.RegisterDockerEventListener(client, generateConfig, &wg, laziness)
	wg.Wait()

	log.Info("done")
}

func generateConfig() {
	log.Info("Triggering refresh...")
	utils.Refresh.Mu.Lock()
	utils.Refresh.IsTriggered = false
	utils.Refresh.Mu.Unlock()
	forwarder.TriggerRefresh(client, getLogstashEndpoint(), configFile)
}

func getDockerEndpoint() string {
	return utils.EndPoint("unix:///var/run/docker.sock", dockerEndPoint, "DOCKER_HOST")
}

func getLogstashEndpoint() string {
	return utils.EndPoint("ec2-54-174-166-239.compute-1.amazonaws.com:5043", logstashEndPoint, "LOGSTASH_HOST")
}
