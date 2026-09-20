package gui

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	Config "github.com/ShuffleKamori/go-practice/config"
	"github.com/ShuffleKamori/go-practice/events"
)

var Window fyne.App

const (
	itemsPerPage = 6
	cols         = 3
	cellW        = 150.0
	cellH        = 250.0
	gridX        = 200.0
	gridY        = 450.0
)

var (
	catalog  []Config.Buycard
	filtered []Config.Buycard
	page     = 0
	area     *fyne.Container
)

func initialize(cfg *Config.WindowApp) fyne.Window {
	Window = app.New()
	mainWindow := Window.NewWindow(cfg.NameWindow)
	if deskWindow, ok := mainWindow.(desktop.Window); ok {
		deskWindow.RequestPosition(cfg.StartPosX, cfg.StartPosY)
	}
	mainWindow.Resize(fyne.NewSize(cfg.PosX, cfg.PosY))
	mainWindow.SetPadded(false)
	mainWindow.SetFixedSize(true)
	mainWindow.Show()

	return mainWindow
}

func openBuyURL() {
	uri, err := url.Parse("https://rmk.stavedu.ru:8010/?year=2026&month=9")
	if err == nil {
		Window.OpenURL(uri)
	}
}

func Tab(mainWindow fyne.Window, cfg *Config.WindowApp) fyne.CanvasObject {

	generalBtn := widget.NewButton("General", func() {
		General(mainWindow, cfg)
	})
	generalBtn.Importance = widget.LowImportance
	generalBtn.Resize(fyne.NewSize(60, 15))
	generalBtn.Move(fyne.NewPos(350, 15))

	supportBtn := widget.NewButton("Support", func() {
		openBuyURL()
	})
	supportBtn.Importance = widget.LowImportance
	supportBtn.Resize(fyne.NewSize(60, 15))
	supportBtn.Move(fyne.NewPos(250, 15))

	clientBtn := widget.NewButton("Client", func() {
		openBuyURL()
	})
	clientBtn.Importance = widget.LowImportance
	clientBtn.Resize(fyne.NewSize(50, 15))
	clientBtn.Move(fyne.NewPos(450, 15))

	button := container.NewWithoutLayout(generalBtn, supportBtn, clientBtn)
	return button
}

func WelcomMassage(mainWindow fyne.Window) fyne.CanvasObject {
	WelcomMassage1 := canvas.NewText("SUPERIOR ACCELERATION", color.RGBA{R: 66, G: 170, B: 255, A: 255})
	WelcomMassage1.TextSize = 24
	WelcomMassage1.Alignment = fyne.TextAlignCenter
	WelcomMassage1.Move(fyne.NewPos(180, 150))
	WelcomMassage1.Resize(fyne.NewSize(400, 40))

	WelcomMassage2 := canvas.NewText("Buy great bikes at a low price", color.RGBA{R: 66, G: 170, B: 255, A: 255})
	WelcomMassage2.TextSize = 16
	WelcomMassage2.Alignment = fyne.TextAlignCenter
	WelcomMassage2.Move(fyne.NewPos(170, 170))
	WelcomMassage2.Resize(fyne.NewSize(400, 40))

	WelcomMassage := container.NewWithoutLayout(WelcomMassage1, WelcomMassage2)
	return WelcomMassage
}

func PhotoBuyCard(mainWindow fyne.Window, cfg *Config.WindowApp) *canvas.Image {
	BuyCard := canvas.NewImageFromImage(cachedImage(cfg.Pathbuycard))
	BuyCard.FillMode = canvas.ImageFillContain
	return BuyCard
}

var decodedImages = map[string]image.Image{}

func cachedImage(path string) image.Image {
	if img, ok := decodedImages[path]; ok {
		return img
	}
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil
	}
	decodedImages[path] = img
	return img
}

func RenderBuyCard(mainWindow fyne.Window, test *Config.Buycard) fyne.CanvasObject {

	test.Photo.Resize(fyne.NewSize(test.SizeX, test.SizeY))
	test.Photo.Move(fyne.NewPos(test.PosX, test.PosY))

	name := canvas.NewText(test.Name, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	name.Move(fyne.NewPos(test.PosX+60, test.PosY+25))
	price := widget.NewLabel(fmt.Sprintf("%d ₽", test.Price))
	price.Move(fyne.NewPos(test.PosX+70, test.PosY+150))

	buycard := container.NewWithoutLayout(test.Photo, name, price)
	return buycard
}

type Card struct {
	widget.BaseWidget
	content fyne.CanvasObject
	onTap   func()
}

func NewCard(content fyne.CanvasObject, onTap func()) *Card {
	c := &Card{content: content, onTap: onTap}
	c.ExtendBaseWidget(c)
	return c
}

func (c *Card) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.content)
}

func (c *Card) Tapped(_ *fyne.PointEvent)          { c.onTap() }
func (c *Card) TappedSecondary(_ *fyne.PointEvent) {}

func CardClick(p Config.Buycard, cfg *Config.WindowApp) {
	detail := Window.NewWindow("Товар: " + p.Name)
	if deskW, ok := detail.(desktop.Window); ok {
		deskW.RequestPosition(cfg.StartPosX+int(cfg.PosX)+5, cfg.StartPosY)
	}
	detail.Resize(fyne.NewSize(320, 460))
	detail.SetFixedSize(true)

	bg := canvas.NewRectangle(cfg.ColorWindow)
	bg.Resize(fyne.NewSize(320, 460))
	bg.Move(fyne.NewPos(0, 0))

	photo := PhotoBuyCard(nil, cfg)
	photo.Resize(fyne.NewSize(250, 250))
	photo.Move(fyne.NewPos(35, 50))

	name := canvas.NewText("Название: "+p.Name, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	name.Move(fyne.NewPos(10, 320))
	name.Resize(fyne.NewSize(300, 30))

	price := canvas.NewText(fmt.Sprintf("Цена: %d ₽", p.Price), color.RGBA{R: 255, G: 255, B: 255, A: 255})
	price.Move(fyne.NewPos(10, 355))
	price.Resize(fyne.NewSize(300, 30))

	buyBtn := widget.NewButton("Купить", func() {
		openBuyURL()
	})
	buyBtn.Resize(fyne.NewSize(200, 40))
	buyBtn.Move(fyne.NewPos(60, 400))

	detail.SetContent(container.NewWithoutLayout(bg, photo, name, price, buyBtn))
	detail.Show()
}

func FiltersPanel(mainWindow fyne.Window, cfg *Config.WindowApp) fyne.CanvasObject {
	nameLabel := canvas.NewText("Имя", color.RGBA{R: 255, G: 255, B: 255, A: 255})
	nameLabel.Move(fyne.NewPos(20, 420))
	nameLabel.Resize(fyne.NewSize(150, 20))

	filterName := widget.NewEntry()
	filterName.Resize(fyne.NewSize(150, 30))
	filterName.Move(fyne.NewPos(20, 440))

	priceLabel := canvas.NewText("Цена", color.RGBA{R: 255, G: 255, B: 255, A: 255})
	priceLabel.Move(fyne.NewPos(20, 480))
	priceLabel.Resize(fyne.NewSize(150, 20))

	filterPrice := widget.NewEntry()
	filterPrice.Resize(fyne.NewSize(150, 30))
	filterPrice.Move(fyne.NewPos(20, 500))

	apply := func() {
		filtered = nil
		for _, p := range catalog {
			ok := true
			if n := strings.TrimSpace(filterName.Text); n != "" &&
				!strings.Contains(strings.ToLower(p.Name), strings.ToLower(n)) {
				ok = false
			}
			if pn := strings.TrimSpace(filterPrice.Text); ok && pn != "" {
				limit, err := strconv.Atoi(pn)
				if err != nil || p.Price > limit {
					ok = false
				}
			}
			if ok {
				filtered = append(filtered, p)
			}
		}
		page = 0
		if area != nil {
			area.Objects[0] = pageContent(mainWindow, cfg)
			area.Refresh()
		}
	}

	var debounce *time.Timer
	applyLater := func() {
		if debounce != nil {
			debounce.Stop()
		}
		debounce = time.AfterFunc(300*time.Millisecond, func() {
			fyne.Do(apply)
		})
	}
	filterName.OnChanged = func(string) { applyLater() }
	filterPrice.OnChanged = func(string) { applyLater() }

	return container.NewWithoutLayout(nameLabel, filterName, priceLabel, filterPrice)
}

func pageContent(mainWindow fyne.Window, cfg *Config.WindowApp) fyne.CanvasObject {
	objs := make([]fyne.CanvasObject, 0, itemsPerPage)
	for i := 0; i < itemsPerPage; i++ {
		idx := page*itemsPerPage + i
		if idx >= len(filtered) {
			break
		}
		b := filtered[idx]
		b.Photo = PhotoBuyCard(mainWindow, cfg)
		b.Photo.Resize(fyne.NewSize(200, 160))
		b.Photo.Move(fyne.NewPos(0, 0))

		name := canvas.NewText(b.Name, color.RGBA{R: 255, G: 255, B: 255, A: 255})
		name.Alignment = fyne.TextAlignCenter
		name.Move(fyne.NewPos(0, 165))
		name.Resize(fyne.NewSize(200, 30))

		price := canvas.NewText(fmt.Sprintf("%d ₽", b.Price), color.RGBA{R: 255, G: 255, B: 255, A: 255})
		price.Alignment = fyne.TextAlignCenter
		price.Move(fyne.NewPos(0, 197))
		price.Resize(fyne.NewSize(200, 26))

		card := NewCard(
			container.NewWithoutLayout(b.Photo, name, price),
			func() { CardClick(b, cfg) },
		)
		card.Move(fyne.NewPos(gridX+float32(i%cols)*cellW, gridY+float32(i/cols)*cellH))
		card.Resize(fyne.NewSize(b.SizeX, 225))
		objs = append(objs, card)
	}
	return container.NewWithoutLayout(objs...)
}

func flip(mainWindow fyne.Window, cfg *Config.WindowApp, delta int) {
	np := page + delta
	if np < 0 || np*itemsPerPage >= len(filtered) {
		return
	}
	page = np
	if area != nil {
		area.Objects[0] = pageContent(mainWindow, cfg)
		area.Refresh()
	}
}

func watchResults(mainWindow fyne.Window, cfg *Config.WindowApp) {
	for {
		if data, ok := events.CheckResult(); ok {
			if products, isCat := data.([]Config.Buycard); isCat {
				catalog = products
				filtered = products
				fyne.Do(func() {
					area.Objects[0] = pageContent(mainWindow, cfg)
					area.Refresh()
				})
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func Pages(mainWindow fyne.Window) fyne.CanvasObject {
	PagesFirst := widget.NewButton("General", func() {
		//
	})
	PagesFirst.Importance = widget.LowImportance
	PagesFirst.Resize(fyne.NewSize(60, 15))
	PagesFirst.Move(fyne.NewPos(350, 15))

	PagesSecond := widget.NewButton("Support", func() {
		//
	})
	PagesSecond.Importance = widget.LowImportance
	PagesSecond.Resize(fyne.NewSize(60, 15))
	PagesSecond.Move(fyne.NewPos(250, 15))

	PagesThird := widget.NewButton("Client", func() {
		//
	})
	PagesThird.Importance = widget.LowImportance
	PagesThird.Resize(fyne.NewSize(50, 15))
	PagesThird.Move(fyne.NewPos(450, 15))

	PagesMax := widget.NewButton(">", func() {
		//
	})
	PagesMax.Importance = widget.LowImportance
	PagesMax.Resize(fyne.NewSize(50, 15))
	PagesMax.Move(fyne.NewPos(450, 15))

	PagesMin := widget.NewButton("<", func() {
		//
	})
	PagesMin.Importance = widget.LowImportance
	PagesMin.Resize(fyne.NewSize(50, 15))
	PagesMin.Move(fyne.NewPos(450, 15))

	button := container.NewWithoutLayout(PagesFirst, PagesSecond, PagesThird, PagesMax, PagesMin)

	return button

}

func General(mainWindow fyne.Window, cfg *Config.WindowApp) {

	backgroundColor := canvas.NewRectangle(cfg.ColorWindow)
	backgroundImage := canvas.NewImageFromFile(cfg.Pathbackground)
	backgroundImage.FillMode = canvas.ImageFillStretch
	backgroundImage.Resize(fyne.NewSize(cfg.PosX, 400))
	imageContainer := container.NewWithoutLayout(backgroundImage)

	tab := Tab(mainWindow, cfg)
	WelcomMassage := WelcomMassage(mainWindow)

	area = container.NewMax(pageContent(mainWindow, cfg))
	go watchResults(mainWindow, cfg)
	events.SendRequest("GET_CATALOG")

	prev := widget.NewButton("<", func() { flip(mainWindow, cfg, -1) })
	prev.Importance = widget.LowImportance
	prev.Resize(fyne.NewSize(30, 25))
	prev.Move(fyne.NewPos(425, 410))

	next := widget.NewButton(">", func() { flip(mainWindow, cfg, +1) })
	next.Importance = widget.LowImportance
	next.Resize(fyne.NewSize(30, 25))
	next.Move(fyne.NewPos(460, 410))

	nav := container.NewWithoutLayout(prev, next)

	filters := FiltersPanel(mainWindow, cfg)

	menu := container.NewMax(backgroundColor, imageContainer, tab, WelcomMassage, filters, area, nav)
	mainWindow.SetContent(menu)
}

func Start_gui() bool {
	cfg := Config.WindowApp{
		ColorWindow:    color.RGBA{R: 25, G: 25, B: 25, A: 255},
		Pathbackground: "background/background1.jpg",
		Pathbuycard:    "BuyCard/Recycle.jpg",
		NameWindow:     "Bicycle store by Shuffle",
		PosX:           750,
		PosY:           1000,
		StartPosX:      300,
		StartPosY:      50,
	}

	mainWindow := initialize(&cfg)
	General(mainWindow, &cfg)

	mainWindow.ShowAndRun()

	return true
}
