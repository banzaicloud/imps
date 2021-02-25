
# Image URL to use all building/pushing image targets
IMG ?= imagepullsecrects:latest
CHART_NAME = imagepullsecrets

include ../service.mk
