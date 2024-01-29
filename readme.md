- window
  - chocolety install
    - powershell 에서 아래 command 실행하여 설치(https://chocolatey.org/install)
      - Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
  - git install
    - choco install git
  - go install 
    - choco install go
  - mariadb install
    - choco install mariadb --version=10.2.14
    - set schema
      ```
      CREATE DATABASE `ab180` /*!40100 COLLATE 'utf8mb4_general_ci' */;
    
      CREATE TABLE `short_link` (
      `id` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
      `created_at` TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
      `url` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
      PRIMARY KEY (`id`) USING BTREE;
      )
      COLLATE='utf8mb4_general_ci'
      ENGINE=InnoDB;
      ``` 
  - influxdb install
    - choco install influxdb2
  - grafana install
    - choco install grafana --version=6.1.6
    ```
      from(bucket: "ab180")
      |> range(start: -7d)
      |> filter(fn: (r) => r["_measurement"] == "data")
      |> aggregateWindow(every: 1d, fn: count)
      |> yield(name: "_value")
    ```

- mac
  - homebrew install
  - go install
  - mariadb install
  - influxdb install
  - grafana install