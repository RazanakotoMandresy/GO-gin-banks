#!/bin/bash
docker buid -t razanakotomandresy/go-gin-banks .
docker push razanakotomandresy/go-gin-banks
docker compose build
docker compose up
