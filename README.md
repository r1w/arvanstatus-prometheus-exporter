arvanstatus-prometheus-exporter




arvanstatus-prometheus-exporter is a Prometheus exporter designed to monitor the status of ArvanCloud services by fetching data from arvancloudstatus.ir. This project is written in Golang and provides metrics that can be scraped by Prometheus to track the uptime and performance of ArvanCloud services.
Table of Contents

    Installation
    Usage
    Features
    Configuration
    Contributing
    License
    Contact

Installation
Prerequisites

    Golang 
    Prometheus

Steps

    Clone the repository:

    bash

git clone https://github.com/r1w/arvanstatus-prometheus-exporter

Navigate into the directory:

bash

cd arvanstatus-prometheus-exporter

Build the project:

bash

go build -o arvanstatus-exporter

Run the exporter:

bash

./arvanstatus-exporter

Add to Prometheus configuration:

Add the following job to your Prometheus prometheus.yml configuration file:

yaml

    scrape_configs:
      - job_name: 'arvanstatus'
        static_configs:
          - targets: ['localhost:your_exporter_port']

    Restart Prometheus:

    Restart Prometheus to apply the new configuration.

Usage

Once the exporter is running, Prometheus will start scraping metrics from it. The exporter listens on a specific port (default: :8080) and exposes metrics that include the status of various ArvanCloud services.

Example metrics include:

plaintext

arvan_service_status{name="cdn"} 1
arvan_service_status{name="dns"} 0

Command-line Flags

    --port: Specify the port on which the exporter listens (default: 8080).
    --interval: Set the scrape interval for fetching ArvanCloud status (default: 60s).

Example:

bash

./arvanstatus-exporter --port 9090 --interval 120s

Features

    Real-time Monitoring: Continuously fetches and exports the status of ArvanCloud services.
    Customizable: Supports configuration of the scrape interval and listening port.
    Lightweight: Minimal resource usage, ideal for deployment alongside other Prometheus exporters.

Configuration

You can configure the exporter via command-line flags or environment variables:

    Command-line Flags:
        --port: Port for the exporter to listen on.
        --interval: Scrape interval for fetching the service status.

    Environment Variables:
        EXPORTER_PORT: Set the port number.
        SCRAPE_INTERVAL: Set the interval for scraping the status.

Example:

bash

export EXPORTER_PORT=9090
export SCRAPE_INTERVAL=120s
./arvanstatus-exporter

Contributing

We welcome contributions to improve the exporter. To contribute:

    Fork the repository
    Create a new branch (git checkout -b feature-branch)
    Make your changes and commit them (git commit -m 'Add some feature')
    Push to the branch (git push origin feature-branch)
    Open a Pull Request

Please follow the coding style and standards outlined in the repository.
License

This project is licensed under the MIT License - see the LICENSE file for details.
Contact

    Your Name - Your Twitter Handle
    Email: kurosch86@gmail.com
    GitHub: https://github.com/r1w/
    Linkedin: https://www.linkedin.com/in/hamid-hadigol/
