/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	os "os"
	"strings"
	"time"
)

var Log *log.Logger

func NewLogger() *log.Logger {
	if Log != nil {
		return Log
	}

	path := "./efa.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	Log = log.New()
	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			log.InfoLevel:  writer,
			log.ErrorLevel: writer,
		},
		&log.JSONFormatter{},
	))
	//log.SetOutput(Log.Writer())
	return Log
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "efa",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	Log.Info("Execute: ", strings.Join(os.Args, " "))
	err := rootCmd.Execute()
	if err != nil {
		Log.Error(err)
		os.Exit(1)
	}
}

func init() {
	NewLogger()
}
