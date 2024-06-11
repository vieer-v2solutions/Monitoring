from flask import Flask, render_template
from prometheus_flask_exporter import PrometheusMetrics
import psutil
import os

app = Flask(__name__)
metrics = PrometheusMetrics(app)

function_calls_counter = metrics.counter(
    'custom_function_called', 'Total number of function calls'
)

memory_usage_gauge = metrics.gauge(
    'memory_usage_bytes', 'Current memory usage in bytes'
)

@app.route("/")
def hello():
    function_calls_counter
    return render_template('index.html')

@metrics.do_not_track()
def update_memory_usage():
    process = psutil.Process(os.getpid())
    memory_usage_gauge.set(process.memory_info().rss)

@app.route('/metrics')
def expose_metrics():
    return metrics.export_http()

if __name__ == "__main__":
    app.run(debug=False, host="0.0.0.0", port=8000)
