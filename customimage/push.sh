#!/bin/bash
exec(docker build -t deepatp/golangcentos:base .)
exec(docker login --username deepatp --password rathna4582)
exec(docker push deepatp/golangcentos:base)
