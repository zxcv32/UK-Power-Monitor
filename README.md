# UK Power Monitor

This app monitors the mains power status of the Uttarakhand state.

The change in mains power status is tweeted using [UK-Power-Monitor-Twitter-Bot](https://github.com/zxcv32/UK-Power-Monitor-Twitter-Bot)

[![Bot Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/UKPowerMonitor.svg?style=social&label=Follow%20%40UKPowerMonitor)](https://twitter.com/UKPowerMonitor)
[![Creator Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/i14a23h19a.svg?style=social&label=Follow%20%40i14a23h19a)](https://twitter.com/i14a23h19a)

Functions: 
1. Monitors mains power and other stats. 
2. Stores metric in a Timeseries database (InfluxDB).
3. Display stats on a Grafana Dashboard.

# Grafana
![img.png](img.png)
[Grafana Dashboard Snapshot](https://snapshots.raintank.io/dashboard/snapshot/EjkITqt85gGjJwfb2tf07lsEDsm3M5hs)

Or import [Grafana Dashboard JSON](./Grafana%20Dashboard.json)

## Components
1. Raspberry Pi 4 Model B
2. Arduino Nano
3. DHT11
4. Voltage Sensor
5. Power backup
6. External storage (optional)

Checkout [Arm-Unit](https://github.com/zxcV32/Arm-Unit) for modular case for this project.

# Deployment
> Deploy Grafana and InfluxDB using [ansible playbooks](https://github.com/zxcV32/ansible-playbooks) 

1. Write `nano/sketch/sketch.ino` to Arduino nano.
  > Install `DHT` and `ArduinoJson` library first
2.Build and deploy docker container from `pi` directory
