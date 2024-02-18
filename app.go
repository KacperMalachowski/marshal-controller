package main

import (
	"changeme/pkg/station"
	"context"
	"crypto/sha256"
	"encoding/json"
	"net"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx         context.Context
	stationHash []byte
	station     station.Definition
	conn        net.Conn
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) shutdown(ctx context.Context) {
}

func (a *App) LoadStationFile() (station.Definition, string, error) {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return a.station, string(a.stationHash), err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return a.station, string(a.stationHash), err
	}

	var station station.Definition
	err = json.Unmarshal(data, &station)
	if err != nil {
		return a.station, string(a.stationHash), err
	}

	a.station = station
	a.stationHash = calculateSHA256(data)

	return station, string(a.stationHash), nil
}

func calculateSHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
