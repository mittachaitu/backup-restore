package signals

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// CancelOnShutdown starts a goroutine that will call cancelFunc when
// either SIGINT or SIGTERM is received
func CancelOnShutdown(cancelFunc context.CancelFunc, logger logrus.FieldLogger) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Infof("Received signal %s, shutting down", sig)
		cancelFunc()
	}()
}
