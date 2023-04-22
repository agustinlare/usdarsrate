# Notifcacion USD to ARS

Manda las cotizaciones de [Dolarito](https://www.dolarito.ar/) por webhook a Google Meets y Discord son los unicos con que probe.

## Envs

*WEBHOOK_URL*: Obviamente el webhook que requiere para mandar la notificaicon
*ENDPOINT_URL*: Por default uso https://www.dolarito.ar/api/frontend/quotations, pero se puede pisar.
*INFLUXDB_URL*: Ademas podes ponerlo en una db de influx para verlo por grafana.

## Docker / Kubernetes

docker run -e WEBHOOK_URL="<WEBHOOK>" -e ENDPOINT_URL="<ENDPOINT>" -e INFLUXDB_URL="<INFLUXDB_URL>" -itd quay.io/agustinlare/usdrate

### Tags
* x86: latest
* Arm: arm

Se puede bajar tal cual esta y compilar con los valores deseados, utilizar el yaml que dejo aca para que corra como cronjob en kubernetes o usar esta imagen `quay.io/agustinlare/dolar-wh`.

## InfluxDB

Conectado al influxdb estos son los comandos para crear la base de datos

```
CREATE DATABASE dolar WITH DURATION 30d REPLICATION 1 SHARD DURATION 1d NAME autogen
CREATE RETENTION POLICY "one_month" ON "dolar" DURATION 30d REPLICATION 1 DEFAULT
```

```
> SELECT "sell" AS "mean_sell" FROM "dolar"."one_month"."exchange_rate"
name: exchange_rate
time                mean_sell
----                ---------
1682193054000000000 442,00
1682193211000000000 442,00
1682193314000000000 442,00
```