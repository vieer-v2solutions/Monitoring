from flask import Flask, request, render_template
from prometheus_flask_exporter import PrometheusMetrics
import time
import os
import psutil

app = Flask(__name__, template_folder='templates')
metrics = PrometheusMetrics(app)

# Custom metric to count function calls
function_calls_counter = metrics.counter(
    'function_calls_total', 'Total number of function calls'
)

# Custom metric to measure memory consumption
memory_usage_gauge = metrics.gauge(
    'memory_usage_bytes', 'Current memory usage in bytes'
)

# Function to be called
def process_number(number):
    # Simulate processing
    time.sleep(1)
    print(f"Processing number {number}")

@app.route('/', methods=['GET'])
def index():
    return render_template('index.html')

# Update memory usage metric periodically
@metrics.do_not_track()
def update_memory_usage():
    process = psutil.Process(os.getpid())
    memory_usage_gauge.set(process.memory_info().rss)

@app.route('/metrics')
def expose_metrics():
    return metrics.export_http()

if __name__ == '__main__':
    metrics.start_http_server(8000)
    app.run(debug=True)
