package wsempty

import (
	"fmt"
	"strings"
	"time"

	"github.com/SpaiR/imgui-go"
	"sdmm/app/ui/cpwsarea/workspace"
	"sdmm/dm/dmenv"
)

type Action interface {
	AppDoOpenEnvironment()
	AppDoOpenEnvironmentByPath(path string)
	AppDoOpenMap()
	AppDoOpenMapByPathV(path string, workspaceIdx int)
	AppHasLoadedEnvironment() bool
	AppRecentEnvironments() []string
	AppRecentMapsByLoadedEnvironment() []string
	AppLoadedEnvironment() *dmenv.Dme
}

type WsEmpty struct {
	workspace.Base

	action Action
	name   string
}

func New(action Action) *WsEmpty {
	name := fmt.Sprint("New##workspace_empty_", time.Now().Nanosecond())
	ws := &WsEmpty{action: action, name: name}
	ws.Workspace = ws
	return ws
}

func (ws *WsEmpty) Name() string {
	return ws.name
}

func (ws *WsEmpty) Process() {
	if !ws.action.AppHasLoadedEnvironment() {
		ws.showEnvironmentsControl()
	} else {
		ws.showMapsControl()
	}
}

func (ws *WsEmpty) showEnvironmentsControl() {
	if imgui.Button("Open Environment...") {
		ws.action.AppDoOpenEnvironment()
	}
	imgui.Separator()
	if len(ws.action.AppRecentEnvironments()) == 0 {
		imgui.Text("No Recent Environments")
	} else {
		imgui.Text("Recent Environments:")
		for _, envPath := range ws.action.AppRecentEnvironments() {
			if imgui.SmallButton(envPath) {
				ws.action.AppDoOpenEnvironmentByPath(envPath)
			}
		}
	}
}

func (ws *WsEmpty) showMapsControl() {
	imgui.Text(fmt.Sprint("Environment: ", ws.action.AppLoadedEnvironment().RootFile))
	imgui.Separator()
	if imgui.Button("Open Map...") {
		ws.action.AppDoOpenMap()
	}
	imgui.Separator()
	if len(ws.action.AppRecentMapsByLoadedEnvironment()) == 0 {
		imgui.Text("No Recent Maps")
	} else {
		imgui.Text("Recent Maps:")
		for _, mapPath := range ws.action.AppRecentMapsByLoadedEnvironment() {
			if imgui.SmallButton(sanitizeMapPath(ws.action.AppLoadedEnvironment().RootDir, mapPath)) {
				ws.action.AppDoOpenMapByPathV(mapPath, ws.Idx())
			}
		}
	}
}

func sanitizeMapPath(envRootDir, mapPath string) string {
	return strings.Replace(mapPath, envRootDir, "", 1)[1:]
}
