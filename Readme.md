# Custom Metrics Setup with Prometheus in Go
## Purposes
The purpose of this documentation is to demonstrate the implementation of custom metrics in a Go application and how to connect them to Prometheus for monitoring.

## Setup Instructions

### Prometheus Configuration

Ensure that the following configurations are present in the `prometheus.yml` file, typically located at `/etc/prometheus/prometheus.yml`:

```yaml
# Prometheus Configuration File
scrape_configs:
  - job_name: 'my_go_app'
    static_configs:
      - targets: ['localhost:9090'] # Change this to your Go application's address
