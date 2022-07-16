# Cotizacion por chat

Manda la cotizacion del dolar a un chat via webhook

## Utilizacion

Le pasas por enviroment un webhook (WEBHOOK_URL) y el endpoint (ENDPOINT_URL) que queres revisar.

```
export WEBHOOK_URL="https://webhook.com/trucho1239801723123"
export ENDPOINT_URL="/api/dolarblue"

# Se pueden pasar separado por coma para obtener varias
export ENDPOINT_URL="/api/dolarblue, /api/dolarbolsa"
```

## Docker / Kubernetes

Se puede bajar tal cual esta y compilar con los valores deseados, utilizar el yaml que dejo aca para que corra como cronjob en kubernetes o usar esta imagen `quay.io/agustinlare/dolarbychat`.

## Endpoints
URL: https://api-dolar-argentina.herokuapp.com/

| Metodo | Endpoint | Descripcion |
| ------ | ------ | ------ |
| GET | /api/dolaroficial | Cotizacion dólar oficial |
| GET | /api/dolarblue | Cotizacion dólar blue |
| GET | /api/contadoliqui | Cotizacion dólar contado con liqui |
| GET | /api/dolarpromedio | Cotizacion dólar promedio |
| GET | /api/dolarturista | Cotizacion dólar turista |
| GET | /api/dolarbolsa | Cotizacion dólar bolsa |
| GET | /api/riesgopais | Valor riesgo pais |
| GET | /api/bbva | Cotizacion dolar del Banco BBVA |
| GET | /api/piano | Cotizacion dolar del Banco Piano |
| GET | /api/hipotecario | Cotizacion dolar del Banco Hipotecario |
| GET | /api/galicia | Cotizacion dolar del Banco Galicia |
| GET | /api/santander | Cotizacion dolar del Banco Santander |
| GET | /api/ciudad | Cotizacion dolar del Banco Ciudad |
| GET | /api/supervielle | Cotizacion dolar del Banco Supervielle |
| GET | /api/patagonia | Cotizacion dolar del Banco Patagonia |
| GET | /api/comafi | Cotizacion dolar del Banco Comafi |
| GET | /api/nacion | Cotizacion dolar del Banco Nación |
| GET | /api/bind | Cotizacion dolar del Banco Industrial |
| GET | /api/bancor | Cotizacion dolar del Banco de Córdoba |
| GET | /api/chaco | Cotizacion dolar del Nuevo Banco del Chaco |
| GET | /api/pampa | Cotizacion dolar del Banco de La Pampa |
| GET | /api/mayorista | Cotizacion dolar Mayorista Bancos|
| GET | /api/euro/nacion | Cotizacion del Euro del Banco Nación |
| GET | /api/euro/galicia | Cotizacion del Euro del Banco Galicia |
| GET | /api/euro/bbva | Cotizacion del Euro del Banco BBVA |
| GET | /api/euro/pampa | Cotizacion del Euro del Banco de La Pampa |
| GET | /api/euro/chaco | Cotizacion del Euro del Nuevo Banco del Chaco |
| GET | /api/euro/hipotecario | Cotizacion del Euro del Banco Hipotecario |
| GET | /api/real/nacion | Cotizacion del Real del Banco Nación |
| GET | /api/real/bbva | Cotizacion del Real del Banco BBVA |
| GET | /api/real/chaco | Cotizacion del Real del Nuevo Banco del Chaco |
| GET | /api/all | Todos los valores disponbiles

## Creditos
Yo casi no hice nada, todo el laburo se lo merce el que creo esta api https://github.com/Castrogiovanni20/api-dolar-argentina/
