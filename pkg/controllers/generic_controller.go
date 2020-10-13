package controllers

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type genericController struct {
	name             string
	queue            workqueue.RateLimitingInterface
	logger           logrus.FieldLogger
	syncHandler      func(obj interface{}) error
	cacheSyncWaiters []cache.InformerSynced
}

func newGenericController(name string, logger logrus.FieldLogger) *genericController {
	c := &genericController{
		name:   name,
		queue:  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), name),
		logger: logger.WithField("controller", name),
	}

	return c
}

// Run is a blocking function that runs the specified number of worker goroutines
// to process items in the work queue. It will return when it receives on the
// ctx.Done() channel.
func (c *genericController) Run(numWorkers int, ctx context.Context) error {
	if c.syncHandler == nil {
		// programmer error
		panic("syncHandler is required to process the request")
	}

	var wg sync.WaitGroup

	defer func() {
		c.logger.Info("Waiting for workers to finish their work")

		c.queue.ShutDown()

		// We have to wait here in the deferred function instead of at the bottom of the function body
		// because we have to shut down the queue in order for the workers to shut down gracefully, and
		// we want to shut down the queue via defer and not at the end of the body.
		wg.Wait()

		c.logger.Info("All workers have finished")

	}()

	c.logger.Info("Starting controller %s", c.name)
	defer c.logger.Info("Shutting down controller %s", c.name)

	// only want to log about cache sync waiters if there are any
	if len(c.cacheSyncWaiters) > 0 {
		c.logger.Info("Waiting for caches to sync")
		if !cache.WaitForCacheSync(ctx.Done(), c.cacheSyncWaiters...) {
			return errors.New("timed out waiting for caches to sync")
		}
		c.logger.Info("Caches are synced")
	}

	if c.syncHandler != nil {
		wg.Add(numWorkers)
		for i := 0; i < numWorkers; i++ {
			go func() {
				wait.Until(c.runWorker, time.Second, ctx.Done())
				wg.Done()
			}()
		}
	}

	<-ctx.Done()

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *genericController) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *genericController) processNextWorkItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// always call done on this item, since if it fails we'll add
	// it back with rate-limiting below
	defer c.queue.Done(key)

	err := c.syncHandler(key)
	if err == nil {
		// If you had no error, tell the queue to stop tracking history for your key. This will reset
		// things like failure counts for per-item rate limiting.
		c.queue.Forget(key)
		return true
	}

	c.logger.WithError(err).WithField("key", key).Error("Error in syncHandler, re-adding item to queue")
	// we had an error processing the item so add it back
	// into the queue for re-processing with rate-limiting
	c.queue.AddRateLimited(key)

	return true
}

func (c *genericController) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		c.logger.WithError(errors.WithStack(err)).
			Error("Error creating queue key, item not added to queue")
		return
	}

	c.queue.Add(key)
}

func (c *genericController) enqueueSecond(_, obj interface{}) {
	c.enqueue(obj)
}
