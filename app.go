package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	goruntime "runtime"

	"github.com/kacpermalachowski/marshal-controller/internal/build"
	"github.com/kacpermalachowski/marshal-controller/pkg/station"
	"github.com/kacpermalachowski/marshal-controller/pkg/td2"

	selfupdate "github.com/creativeprojects/go-selfupdate"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx         context.Context
	stationHash []byte
	station     station.Definition
	client      *td2.Client
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.LogInfo(ctx, fmt.Sprintf("Version %s, Time %s", build.Version, build.Time))
}

func (a *App) domReady(ctx context.Context) {
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) shutdown(ctx context.Context) {
}

func (a *App) LoadStationFile() station.Definition {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprint(err))
		return a.station
	}

	data, err := os.ReadFile(file)
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprint(err))
		return a.station
	}

	station, err := station.ParseStationDefinition(data)
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprint(err))
		return a.station
	}

	a.station = station
	a.stationHash = calculateSHA256(data)
	a.client = td2.New(a.ctx, fmt.Sprintf("%x", a.stationHash))
	runtime.LogInfof(a.ctx, "Loaded station file with hash: %s", a.stationHash)

	return station
}

func (a *App) Connect(address string) error {
	err := a.client.Connect(address)
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprint(err, address))
		return err
	}

	go func() {
		for {
			message := <-a.client.ReadChan

			if message == "" {
				continue
			}

			runtime.LogInfo(a.ctx, fmt.Sprint(message))
			runtime.EventsEmit(a.ctx, "message", message)
		}
	}()

	return nil
}

func (a *App) Disconnect() {
	a.client.Disconnect()
}

func (a *App) SetSignal(hill station.Hill, signal string) {
	if !a.client.IsConnected {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Disconnected",
			Message: "You cannot set signal while you're disconnected!",
		})
		return
	}

	runtime.LogInfo(a.ctx, fmt.Sprint("Setting: ", signal))
	err := a.client.Write(fmt.Sprintf("%s:%s", hill.Signal, signal))
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprint(err))
		return
	}

	for _, repeater := range hill.Repeaters {
		err := a.client.Write(fmt.Sprintf("%s:%s", repeater, signal))
		if err != nil {
			runtime.LogError(a.ctx, fmt.Sprint(err))
			return
		}
	}
}

func (a *App) GetStationHash() string {
	return fmt.Sprintf("%x", a.stationHash)
}

func (a *App) CheckForUpdate(currentVersion string) (bool, error) {
	latest, found, err := selfupdate.DetectLatest(a.ctx, selfupdate.ParseSlug("KacperMalachowski/marshal-controller"))
	if err != nil {
		return false, fmt.Errorf("error occurred while detecting version: %w", err)
	}
	if !found {
		return false, fmt.Errorf("latest version for %s/%s could not be found from github repository", goruntime.GOOS, goruntime.GOARCH)
	}

	runtime.LogDebugf(a.ctx, "Current version %s is less or equal of found %s: %v", latest.Name, currentVersion, latest.LessOrEqual(currentVersion))
	if latest.LessOrEqual(currentVersion) {
		runtime.LogInfof(a.ctx, "Current version (%s) is the latest", currentVersion)
		return false, nil
	}

	return true, nil
}

func calculateSHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
