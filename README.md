# Teonet fortune single-page application (SPA)

This is simple [Teonet](https://github.com/teonet-go/teonet) one-page application which get fortune message from [Teonet Fortune](https://github.com/teonet-go/teofortune) microservice using Google Cloud Function.

[![GoDoc](https://godoc.org/github.com/teonet-go/teofortune-spa?status.svg)](https://godoc.org/github.com/teonet-go/teofortune-spa/)
[![Go Report Card](https://goreportcard.com/badge/github.com/teonet-go/teofortune-spa)](https://goreportcard.com/report/github.com/teonet-go/teofortune-spa)

## Create default frontend

    sudo npm i -g @vue/cli
    vue create frontend

## How to use

To run this application local use next commands:

    go run .

or

    cd frontend
    npm run serve

By default the teofortune-spa site start at localhost:8082. You can publish
this site to Google [Cloud Run](https://console.cloud.google.com/run).

There is preinstalled teofortune-spa web-site in Google Cloud:
<https://fortune-spa.teonet.app>

## Licence

[BSD](LICENSE)
