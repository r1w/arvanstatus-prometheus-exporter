Arvanstatus Prometheus Exporter
=================

Arvanstatus Prometheus Exporter is a Prometheus exporter designed to monitor the status of ArvanCloud services by fetching data from arvancloudstatus.ir. This project is written in Golang and provides metrics that can be scraped by Prometheus to track the uptime and performance of ArvanCloud services.



Table of Contents
=================
<!--ts-->
* [Installation](#installation)
* [Usage](#usage)
* [Features](#Features)
* [Configuration](#Configuration)
* [Contributing](#Contributing)
* [License](#License)
* [Contact](#Contact)
<!--te-->

Installation
--------------

Linux (manual installation)

Clone the repository:
```bash
$ git clone https://github.com/r1w/arvanstatus-prometheus-exporter
```
Navigate into the directory:
```bash
$ cd arvanstatus-prometheus-exporter
```
Build the project:
```bash
$ go build -o arvanstatus-exporter
```
Run the exporter:
```bash
$ ./arvanstatus-exporter
```

Add to Prometheus configuration:

Add the following job to your Prometheus `prometheus.yml` configuration file:

yaml
  ```yaml
    scrape_configs:
      - job_name: 'arvanstatus'
        static_configs:
          - targets: ['localhost:your_exporter_port']
  ```

Restart Prometheus:

    Restart Prometheus to apply the new configuration.

Usage
============
Once the exporter is running, Prometheus will start scraping metrics from it. The exporter listens on a specific port (default: :8080) and exposes metrics that include the status of various ArvanCloud services.

Example metrics include:

```
    arvan_services_status{category="APIs",service="CDN"} 1
    arvan_services_status{category="APIs",service="Cloud Server"} 1
    arvan_services_status{category="APIs",service="Container Service"} 1
    arvan_services_status{category="APIs",service="Edge Computing"} 1
    arvan_services_status{category="APIs",service="Live Streaming"} 1
    arvan_services_status{category="APIs",service="Object Storage"} 1
    arvan_services_status{category="APIs",service="Panel (Middleware)"} 1
    arvan_services_status{category="APIs",service="Video Advertising"} 1
    arvan_services_status{category="APIs",service="Video on-demand Service"} 1
    arvan_services_status{category="CDN",service="CDN"} 1
    arvan_services_status{category="CDN",service="Cloud Security"} 1
    arvan_services_status{category="CDN",service="DNS"} 1
    arvan_services_status{category="CDN PoP-Sites",service="Asia"} 1
    arvan_services_status{category="CDN PoP-Sites",service="Europe"} 1
    arvan_services_status{category="CDN PoP-Sites",service="Iran"} 1
    arvan_services_status{category="CDN PoP-Sites",service="North America"} 1
    arvan_services_status{category="Cloud Server",service="Bamdad (Iran-TehranWest)"} 1
    arvan_services_status{category="Cloud Server",service="Forough (Iran-TehranCentral 2)"} 1
    arvan_services_status{category="Cloud Server",service="Shahriar (Iran-Tabriz)"} 1
    arvan_services_status{category="Cloud Server",service="Simin (Iran-TehranCentral 1)"} 1
    arvan_services_status{category="Container Service",service="Bamdad (Iran-TehranWest)"} 1
    arvan_services_status{category="Container Service",service="Shahriar (Iran-Tabriz)"} 1
    arvan_services_status{category="Managed Database",service="Automatic Failover"} 1
    arvan_services_status{category="Managed Database",service="Automatic Recovery"} 1
    arvan_services_status{category="Managed Database",service="Backup"} 1
    arvan_services_status{category="Managed Database",service="Bamdad"} 1
    arvan_services_status{category="Managed Database",service="MySQL"} 1
    arvan_services_status{category="Managed Database",service="Panel"} 1
    arvan_services_status{category="Managed Database",service="Shahriar (Iran-Tabriz)"} 1
    arvan_services_status{category="Object Storage",service="Shahriar (Iran-Tabriz)"} 1
    arvan_services_status{category="Object Storage",service="Simin (Iran-TehranCentral 1)"} 1
    arvan_services_status{category="Service",service="Radar"} 1
    arvan_services_status{category="Service",service="Website"} 1
    arvan_services_status{category="Support",service="Forum"} 1
    arvan_services_status{category="Support",service="Telephone"} 1
    arvan_services_status{category="Support",service="Ticketing"} 1
    arvan_services_status{category="Video Platform",service="Audio on-demand Service"} 1
    arvan_services_status{category="Video Platform",service="Live Streaming"} 1
    arvan_services_status{category="Video Platform",service="Video Ads"} 1
    arvan_services_status{category="Video Platform",service="Video Player"} 1
    arvan_services_status{category="Video Platform",service="Video on-demand Service"} 1
```


Command-line Flags

    --port: Specify the port on which the exporter listens (default: 8080).
    --interval: Set the scrape interval for fetching ArvanCloud status (default: 60s).

Example:
```bash
$ ./arvanstatus-exporter --port 9090 --interval 120s
```

Features
============

    Real-time Monitoring: Continuously fetches and exports the status of ArvanCloud services.
    Customizable: Supports configuration of the scrape interval and listening port.
    Lightweight: Minimal resource usage, ideal for deployment alongside other Prometheus exporters.

Configuration
============

You can configure the exporter via command-line flags or environment variables:

    Command-line Flags:
        --port: Port for the exporter to listen on.
        --interval: Scrape interval for fetching the service status.

    Environment Variables:
        EXPORTER_PORT: Set the port number.
        SCRAPE_INTERVAL: Set the interval for scraping the status.

Example:

```bash
$ export EXPORTER_PORT=9090
$ export SCRAPE_INTERVAL=120s
$ ./arvanstatus-exporter
```

Contributing
============

We welcome contributions to improve this project. 

To contribute:

    Fork the repository
    Create a new branch (git checkout -b feature-branch)
    Make your changes and commit them (git commit -m 'Add some feature')
    Push to the branch (git push origin feature-branch)
    Open a Pull Request

> [!IMPORTANT]
> Please follow the coding style and standards outlined in the repository.

License:
============

 This project is licensed under the MIT License - see the LICENSE file for details.

Contact:
============
    Your Name - Your Twitter Handle
    Email: kurosch86@gmail.com
    GitHub: https://github.com/r1w/
    Linkedin: https://www.linkedin.com/in/hamid-hadigol/
