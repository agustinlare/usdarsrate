# Cotizacion por chat

Manda la cotizacion del dolar a un chat via webhook

## Utilizacion

Le pasas por enviroment un webhook (WEBHOOK_URL) y el endpoint (ENDPOINT_URL) que queres revisar.

```
export WEBHOOK_URL="https://slack.googlchat.webhook.com"
export ENDPOINT_URL="/api/dolarblue"
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
| GET | /api/casas/caba/metropolis | Cotizacion dolar de Metrópolis Compañía Financiera S.A |
| GET | /api/casas/caba/alpe | Cotizacion dolar de Cambio Alpe S.A |
| GET | /api/casas/caba/vaccaro | Cotizacion dolar de Francisco Vaccaro S.A |
| GET | /api/casas/caba/transatlantica | Cotizacion dolar de Transatlantica S.A |
| GET | /api/casas/lp/tilos | Cotizacion dolar de Casa de Cambio Los Tilos S.A |
| GET | /api/casas/mdp/jonestur | Cotizacion dolar de Jonestur S.A |
| GET | /api/casas/mdp/moneta | Cotizacion dolar de La Moneta Cambio S.A|
| GET | /api/casas/mdp/exchange | Cotizacion dolar de Mar del Plata Exchange S.A |
| GET | /api/casas/rosario/daminato | Cotizacion dolar de Daminato Viajes y Cambios S.A. |
| GET | /api/casas/mendoza/montemar | Cotizacion dolar de Montemar Compañía Financiera S.A |
| GET | /api/casas/tucuman/maguitur | Cotizacion dolar de Maguitur S.A |
| GET | /api/euro/nacion | Cotizacion del Euro del Banco Nación |
| GET | /api/euro/galicia | Cotizacion del Euro del Banco Galicia |
| GET | /api/euro/bbva | Cotizacion del Euro del Banco BBVA |
| GET | /api/euro/pampa | Cotizacion del Euro del Banco de La Pampa |
| GET | /api/euro/chaco | Cotizacion del Euro del Nuevo Banco del Chaco |
| GET | /api/euro/hipotecario | Cotizacion del Euro del Banco Hipotecario |
| GET | /api/real/nacion | Cotizacion del Real del Banco Nación |
| GET | /api/real/bbva | Cotizacion del Real del Banco BBVA |
| GET | /api/real/chaco | Cotizacion del Real del Nuevo Banco del Chaco |
| GET | /api/bcra/reservas | Reservas del BCRA en dólares |
| GET | /api/bcra/circulante | Total de pesos en circulación |
| GET | /api/evolucion/dolaroficial | Evolución mensual anualizada del dólar oficial |
| GET | /api/evolucion/dolarblue | Evolución mensual anualizada del dólar blue |
| GET | /api/evolucion/realoficial | Evolución mensual anualizada del Real oficial |
| GET | /api/evolucion/eurooficial | Evolución mensual anualizada del Euro oficial |
| GET | /api/all | Todos los valores disponbiles

## Creditos
Yo casi no hice nada, todo el laburo se lo merce el que creo esta api https://github.com/Castrogiovanni20/api-dolar-argentina/