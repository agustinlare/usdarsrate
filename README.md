# Cotizacion Webhook Sender

Manda las cotizaciones de [Dolarito](https://www.dolarito.ar/) por webhook.

## Utilizacion

Le pasas por env vars el webhook (WEBHOOK_URL) y el endpoint (ENDPOINT_URL).

```
"env": {
  "WEBHOOK_URL": "https://chat.googleapis.com/v1/webhooktruchisimo=jd645lCtCHp42345bq62wvP345K34BF4RU0FfG9R29khfCIEwM_8-280%3D",
  "ENDPOINT_URL": "https://www.dolarito.ar/api/frontend/quotations"
}
```

## Docker / Kubernetes

Se puede bajar tal cual esta y compilar con los valores deseados, utilizar el yaml que dejo aca para que corra como cronjob en kubernetes o usar esta imagen `quay.io/agustinlare/dolar-wh`.
