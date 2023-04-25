package conf

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	DIALECT  string
	HOST     string
	DBPORT   string
	USER     string
	NAME     string
	PASSWORD string
	ACCESS   string
	REFRESH  string
}

func New() *Conf {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading env file", err)
	}

	conf := &Conf{}
	flag.StringVar(&conf.DIALECT, "DIALECT", os.Getenv("DIALECT"), "dialect")
	flag.StringVar(&conf.HOST, "HOST", os.Getenv("HOST"), "host")
	flag.StringVar(&conf.DBPORT, "DBPORT", os.Getenv("DBPORT"), "port")
	flag.StringVar(&conf.USER, "USER", os.Getenv("USER"), "user")
	flag.StringVar(&conf.NAME, "NAME", os.Getenv("NAME"), "name")
	flag.StringVar(&conf.PASSWORD, "PASSWORD", os.Getenv("PASSWORD"), "password")
	flag.StringVar(&conf.ACCESS, "ACCESS", os.Getenv("ACCESS"), "JWT access key")
	flag.StringVar(&conf.REFRESH, "REFRESH", os.Getenv("REFRESH"), "JWT refresh key")

	return conf
}

func (c *Conf) ConnStr() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		c.HOST, c.USER, c.NAME, c.PASSWORD, c.DBPORT)
}
