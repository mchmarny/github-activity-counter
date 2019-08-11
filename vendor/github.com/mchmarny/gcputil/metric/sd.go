package metric

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3"
	googlepb "github.com/golang/protobuf/ptypes/timestamp"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
	monitoredrespb "google.golang.org/genproto/googleapis/api/monitoredres"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"

	"github.com/mchmarny/gcputil/project"
)

const (
	metricTypePrefix = "custom.googleapis.com/metric"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

// Client represents metric client
type Client struct {
	projectID    string
	metricClient *monitoring.MetricClient
}

// NewClient instantiates client
func NewClient(ctx context.Context) (client *Client, err error) {

	// get project ID
	p, err := project.GetID()
	if err != nil {
		logger.Printf("Error while getting project ID: %v", err)
		return nil, err
	}

	// create metric client
	mc, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		logger.Fatalf("Error creating metric client: %v", err)
	}

	return &Client{
		projectID:    p,
		metricClient: mc,
	}, nil

}

// Publish publishes time series based on metric and value to Stackdriver
// Example: `Publish(ctx, "device1", "friction", 0.125)``
func (c *Client) Publish(ctx context.Context, sourceID, metricType string, metricValue interface{}) error {

	// derive typed value from passed interface
	var val *monitoringpb.TypedValue
	switch v := metricValue.(type) {
	default:
		return fmt.Errorf("Unsupported metric type: %T", v)
	case float64:
		val = &monitoringpb.TypedValue{
			Value: &monitoringpb.TypedValue_DoubleValue{DoubleValue: metricValue.(float64)},
		}
	case int64:
		val = &monitoringpb.TypedValue{
			Value: &monitoringpb.TypedValue_Int64Value{Int64Value: metricValue.(int64)},
		}
	}

	// create data point
	ptTs := &googlepb.Timestamp{Seconds: time.Now().Unix()}
	dataPoint := &monitoringpb.Point{
		Interval: &monitoringpb.TimeInterval{StartTime: ptTs, EndTime: ptTs},
		Value:    val,
	}

	// create time series request with the data point
	tsRequest := &monitoringpb.CreateTimeSeriesRequest{
		Name: monitoring.MetricProjectPath(c.projectID),
		TimeSeries: []*monitoringpb.TimeSeries{
			{
				Metric: &metricpb.Metric{
					Type: fmt.Sprintf("%s/%s", metricTypePrefix, metricType),
					Labels: map[string]string{
						"source_id": sourceID,
						// random label to work around SD complaining
						// about multiple events for same time window
						"rnd_label": fmt.Sprint(rand.Intn(100)),
					},
				},
				Resource: &monitoredrespb.MonitoredResource{
					Type:   "global",
					Labels: map[string]string{"project_id": c.projectID},
				},
				Points: []*monitoringpb.Point{dataPoint},
			},
		},
	}

	// publish series
	return c.metricClient.CreateTimeSeries(ctx, tsRequest)
}
