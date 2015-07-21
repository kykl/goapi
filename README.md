# goapi

## Build
./build.sh <prod|dev|staging>

## Run
docker run -d -P goapi

## Config
Create service account on Google Compute Platform and download *service-account.(json|p12)

## Running locally with Google Cloud
gcloud preview  --verbosity debug  app run ./app.yaml

## Deploying with Google Cloud
gcloud preview  --verbosity debug  app deploy ./app.yaml
