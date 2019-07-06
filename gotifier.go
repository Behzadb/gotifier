package gotifier

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/beeep"
	"github.com/cron"
	"github.com/yaml"
)

// Conf contains api information
type conf struct {
	InfoIcon    string   `yaml:"infoIcon"`
	WarningIcon string   `yaml:"warningIcon"`
	ErrorIcon   []string `yaml:"errorIcon"`
	conf        []struct {
		Name      string   `yaml:"name"`
		Baseurl   string   `yaml:"baseurl"`
		Endpoints []string `yaml:"endpoints"`
		Username  string   `yaml:"username"`
		Password  string   `yaml:"password"`
		Token     string   `yaml:"token"`
		Title     string   `yaml:"title"`
		Fields    []struct {
			ID        string    `yaml:"id"`
			Message   string    `yaml:"message"`
			Timestamp time.Time `yaml:"timestamp"`
			Priority  string    `yaml:"priority"`
		} `yaml:"fields"`
	}
}

// notifications keeps list of notifications
type notifications []struct {
	total        int
	title        string
	notification []struct {
		id        string
		message   string
		appicon   string
		priority  int
		repeat    bool
		link      string
		timestamp time.Time
	}
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// FetchNotif populates notification list
func FetchNotif() {

}

// InitCron schedules a cron job to fetch notifications every "period" minutes
func InitCron(period int) {

	d := cron.New()
	d.AddFunc("@every "+string(period)+"m", func() {
		FetchNotif()
	})
	d.Start()
}

// InvokeNotif invoke desktop notification with parameters
func InvokeNotif(title string, message string, appicon string) {
	err := beeep.Notify(title, message, appicon)
	if err != nil {
		panic(err)
	}
}
