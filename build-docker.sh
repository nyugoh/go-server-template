#!/bin/bash
docker build -t $APP_NAME .
## tag the image
docker tag $APP_NAME
