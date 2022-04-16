# UK Power Monitor

Monitors and tweets mains power status of the Uttarakhand state.

Functions: 
1. Monitors mains power and other stats. 
2. Stores metric in a Timeseries database.
3. Tweet power status (TODO)

# Grafana
![img.png](img.png)
[Grafana Dashboard Snapshot](https://snapshots.raintank.io/dashboard/snapshot/EjkITqt85gGjJwfb2tf07lsEDsm3M5hs)

## Components
1. Raspberry Pi 4 Model B
2. Arduino Nano
3. DHT11
4. Voltage Sensor
5. Power backup
6. External storage (optional)

Checkout [Arm-Unit](https://github.com/zxcV32/Arm-Unit) for modular case for the this project.

# Deployment
> Deploy Grafana and InfluxDB using [ansible playbooks](https://github.com/zxcV32/ansible-playbooks) 

1. Write `nano/sketch/sketch.ino` to Arduino nano.
  > Install `DHT` and `ArduinoJson` library first
2. Build and deploy docker container from `pi` directory
