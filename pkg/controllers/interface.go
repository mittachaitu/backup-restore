package controllers

import "context"

// Interface holds the methods to run the controllers
type Interface interface {
	// Run will set up the event handlers for types we are interested in, as well
	// as syncing informer caches and starting workers. It will block until channel
	// is closed, at which point it will shutdown the workqueue and wait for
	// workers to finish processing their current work items.
	Run(workers int, context context.Context) error
}
