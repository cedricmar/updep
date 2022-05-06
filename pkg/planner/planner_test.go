package planner

import (
	"os"
	"testing"

	"github.com/cedricmar/updep/pkg/depgraph"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCircularDependencies(t *testing.T) {
	gr, err := depgraph.LoadDependencyGraphFromFixture("simple")
	require.NoError(t, err)

	_, err = Plan("events@0.0.1", gr)
	assert.Error(t, err)
}

func TestUpdateDepOnDifferentDepths(t *testing.T) {
	gr, err := depgraph.LoadDependencyGraphFromFixture("ecommerce-app")
	require.NoError(t, err)

	p, err := Plan("event-producer@0.0.1", gr)
	require.NoError(t, err)

	Fprint(os.Stdout, p)

	assert.Len(t, p, 3)

	// STEP 1
	//   event-consumer
	//   control-plane-adapter
	//   auth-adapter
	assert.Len(t, p[0], 3)
	assert.Contains(t, p[0], "event-consumer")
	assert.Contains(t, p[0], "control-plane-adapter")
	assert.Contains(t, p[0], "auth-adapter")

	// STEP 2
	//   key-value-storer
	//   data-planner
	//   jobs-scheduler
	//   user-events-consumer
	//   failure-recovery-actor
	//   data-change-aggregator
	//   messaging-consumer
	//   stock-updates-consumer
	//   inventory-events-consumer
	//   payment-events-consumer
	//   storage-status
	//   gw-events-consumer
	//   k8s-notifications-consumer
	//   credentials-provider
	assert.Len(t, p[1], 14)
	assert.Contains(t, p[1], "key-value-storer")
	assert.Contains(t, p[1], "data-planner")
	assert.Contains(t, p[1], "jobs-scheduler")
	assert.Contains(t, p[1], "user-events-consumer")
	assert.Contains(t, p[1], "failure-recovery-actor")
	assert.Contains(t, p[1], "data-change-aggregator")
	assert.Contains(t, p[1], "messaging-consumer")
	assert.Contains(t, p[1], "stock-updates-consumer")
	assert.Contains(t, p[1], "inventory-events-consumer")
	assert.Contains(t, p[1], "payment-events-consumer")
	assert.Contains(t, p[1], "storage-status")
	assert.Contains(t, p[1], "gw-events-consumer")
	assert.Contains(t, p[1], "k8s-notifications-consumer")
	assert.Contains(t, p[1], "credentials-provider")

	// STEP 3
	//   fraud-prevention-actor
	//   checkout-producer
	assert.Len(t, p[2], 2)
	assert.Contains(t, p[2], "fraud-prevention-actor")
	assert.Contains(t, p[2], "checkout-producer")
}
