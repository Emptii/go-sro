package viewer

import (
	"fmt"
	"github.com/Emptii/go-sro/agent-server/navmeshv2"
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	ToolbarHeight = 64
)

var showTerrainCells = true
var showOnlyBlockedRegionCells = true
var showObjectsInternalEdges = true
var showObjectsGlobalEdges = true

func drawToolbar() {
	toolbarWidth := int32(rl.GetScreenWidth())
	toolbarHeight := int32(ToolbarHeight)
	rl.DrawLine(0, toolbarHeight, toolbarWidth, toolbarHeight, rl.DarkGray)

	const buttonHeight = 40
	const buttonSpacing = 15

	buttonRect := func(x, width float32) rl.Rectangle {
		return rl.Rectangle{
			X:      x,
			Y:      float32(toolbarHeight/2) - float32(buttonHeight/2),
			Width:  width,
			Height: float32(buttonHeight),
		}
	}

	checkBoxRect := func(x float32) rl.Rectangle {
		return rl.Rectangle{
			X:      x,
			Y:      float32(toolbarHeight/2) - float32(buttonHeight/2),
			Width:  float32(buttonHeight),
			Height: float32(buttonHeight),
		}
	}

	doToolbarButton := func(text, description string, bounds rl.Rectangle, onClick func()) (nextX float32) {
		if raygui.Button(bounds, text) {
			onClick()
		}

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), bounds) {
			rl.DrawTextEx(rl.GetFontDefault(), description, rl.NewVector2(20, float32(toolbarHeight)+20), 24, 2, raygui.TextColor())
		}

		return bounds.X + bounds.Width + buttonSpacing
	}

	doCheckbox := func(text string, checked *bool, bounds rl.Rectangle) (nextX float32) {
		*checked = raygui.CheckBox(bounds, *checked)
		textWidth := rl.MeasureText(text, int32(raygui.GetStyleProperty(raygui.GlobalTextFontsize)))
		labelBounds := rl.NewRectangle(bounds.X+bounds.Width+buttonSpacing, bounds.Y, float32(textWidth), bounds.Height)
		raygui.Label(labelBounds, text)

		return labelBounds.X + labelBounds.Width + buttonSpacing
	}

	doCollisionPointBox := func(text string, x, y, z *int, xT, yT, zT *string, currentX float32) (nextX float32) {
		*xT = raygui.TextBox(rl.NewRectangle(currentX, 2, 100, 20), *xT)
		*yT = raygui.TextBox(rl.NewRectangle(currentX, 22, 100, 20), *yT)
		*zT = raygui.TextBox(rl.NewRectangle(currentX, 42, 100, 20), *zT)
		textWidth := rl.MeasureText(text+" X", int32(raygui.GetStyleProperty(raygui.GlobalTextFontsize)))
		raygui.Label(rl.NewRectangle(currentX+100+buttonSpacing, 2, float32(textWidth), 20), text+" X")
		raygui.Label(rl.NewRectangle(currentX+100+buttonSpacing, 22, float32(textWidth), 20), text+" Y")
		raygui.Label(rl.NewRectangle(currentX+100+buttonSpacing, 42, float32(textWidth), 20), text+" Z")

		*x = parseInt(*xT)
		*y = parseInt(*yT)
		*z = parseInt(*zT)

		return currentX + 100 + buttonSpacing + float32(textWidth) + buttonSpacing
	}

	var nextX float32 = buttonSpacing

	nextX = doToolbarButton("Open", "Open PK2 File", buttonRect(nextX, 100), func() {
		loader = navmeshv2.NewLoader("/Users/rmu/workspaces/private/go-sro/agent-server/Data.pk2")
		loader.LoadNavMeshInfos()
		logrus.Info("finished loading navmesh infos")
		loader.LoadTerrainMesh(loader.NavMeshPath+string(os.PathSeparator)+"nv_6587.nvm", 25991)
		loader.LoadTerrainMesh(loader.NavMeshPath+string(os.PathSeparator)+"nv_6588.nvm", 25992)
		loader.LoadTerrainMesh(loader.NavMeshPath+string(os.PathSeparator)+"nv_6687.nvm", 26247)
		//loader.LoadTerrainMesh(loader.NavMeshPath+string(os.PathSeparator)+"nv_6687.nvm", 26246)
		loader.LoadTerrainMesh(loader.NavMeshPath+string(os.PathSeparator)+"nv_61a5.nvm", 24997)
		loader.LoadTerrainMesh(loader.NavMeshPath+string(os.PathSeparator)+"nv_61a6.nvm", 24998)
		loaderProgress = 1.0

		logrus.Info("started loading navmeshes")
		//go loader.LoadTerrainMeshes(loaderProgressChannel)
		//loadingMeshes = true
		//loaderProgressChannel <- loader.MapProjectInfo.ActiveRegionsCount
		//loaderProgressAbsolute = loader.MapProjectInfo.ActiveRegionsCount
		terrainList = loader.RegionData
		logrus.Info("Open File")
	})

	nextX = doToolbarButton("Save as JSON", "Save as JSON", buttonRect(nextX, 100), func() {
		logrus.Info("Saving Navmesh as JSON")
		terrain := terrainList[int16(selectedRegionId)]
		ioutil.WriteFile(fmt.Sprintf("%d.obj", selectedRegionId), terrain.ToBlenderObj(), fs.ModePerm)
		ioutil.WriteFile(fmt.Sprintf("navmesh-viewer/%d.json", selectedRegionId), terrain.ToJson(), fs.ModePerm)
	})
	nextX = doCheckbox("Show Only Blocked Terrain Cells", &showOnlyBlockedRegionCells, checkBoxRect(nextX))
	nextX = doCheckbox("Show Object's Internal Edges", &showObjectsInternalEdges, checkBoxRect(nextX))
	nextX = doCheckbox("Show Object's Global Edges", &showObjectsGlobalEdges, checkBoxRect(nextX))
	nextX = doCheckbox("Draw Objects with 0 height", &drawObjectsWithZeroHeight, checkBoxRect(nextX))

	nextX = doToolbarButton("Reset Camera", "Resets the camera", buttonRect(nextX, 100), func() {
		cam = defaultCam()
		rl.SetCameraMode(cam, rl.CameraFree)
		rl.UpdateCamera(&cam)
	})

	nextX = doCollisionPointBox("P1",
		&collisionPoint1X,
		&collisionPoint1Y,
		&collisionPoint1Z,
		&collisionPoint1XText,
		&collisionPoint1YText,
		&collisionPoint1ZText,
		nextX)
	nextX = doCollisionPointBox("P2",
		&collisionPoint2X,
		&collisionPoint2Y,
		&collisionPoint2Z,
		&collisionPoint2XText,
		&collisionPoint2YText,
		&collisionPoint2ZText,
		nextX)

	raygui.Label(rl.NewRectangle(nextX, 2, 200, 20), fmt.Sprintf("Cells on path: %d", cellsOnPath))
	raygui.Label(rl.NewRectangle(nextX, 22, 200, 20), fmt.Sprintf("Objects on path: %d", objectsOnPath))
	raygui.Label(rl.NewRectangle(nextX, 42, 200, 20), fmt.Sprintf("Object Collisions on path: %d", objectCollisionsOnPath))
	nextX += 200 + buttonSpacing
}

func parseFloat(s string) float32 {
	num, _ := strconv.ParseFloat(s, 32)
	return float32(num)
}

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}
