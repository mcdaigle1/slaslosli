package main

import (
	"log/slog"
	"net/http"

	"github.com/mcdaigle1/slaslosli/config"
	"github.com/mcdaigle1/slaslosli/internal/api"
	"github.com/mcdaigle1/slaslosli/modules/logutils"
    "github.com/mcdaigle1/slaslosli/modules/awsutils"
)

func main() {
    config.Load()

    logutils.InitLogger()

    slog.Debug("loading AWS Secrets")
    err := awsutils.LoadSecrets()
    if err != nil {
        slog.Error("failed to retrieve AWS secrets", "error", err)
    }
    slog.Debug("Prometheus user: " + awsutils.Secrets.PrometheusUsername)

    router := api.NewRouter()

    slog.Info("Starting server on :8080")
    err = http.ListenAndServe(":8080", router)
    if err != nil {
        slog.Error("server failed", "error", err)
    }
}
