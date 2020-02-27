package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangxusc/event-center/pkg/config"
	"github.com/tangxusc/event-center/pkg/receiver"
	"github.com/tangxusc/event-center/pkg/repository"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	newCommand := NewCommand(ctx)
	HandlerNotify(cancel)

	_ = newCommand.Execute()
	cancel()
}

func NewCommand(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:   "start",
		Short: "start server",
		RunE: func(cmd *cobra.Command, args []string) error {
			rand.Seed(time.Now().Unix())
			config.InitLog()
			e := repository.Conn(ctx)
			if e != nil {
				return e
			}
			defer repository.Close(ctx)
			e = receiver.Start(ctx)
			if e != nil {
				return e
			}

			<-ctx.Done()
			return nil
		},
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	config.BindParameter(command)

	return command
}

func HandlerNotify(cancel context.CancelFunc) {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, os.Kill)
		<-signals
		cancel()
	}()
}
