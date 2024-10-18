#!/bin/bash
docker buid -t razanakotomandresy/go-bank-backend 
docker compose build
docker compose up
