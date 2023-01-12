package main 


import (

	"image/color"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/data/binding"

) 


func mass_button_group(str binding.String) fyne.CanvasObject { 

	mass_label := widget.NewLabel("Mass 1:") 
	mass_data := widget.NewLabelWithData(str) 

	mass_entry := widget.NewEntry() 

	mass_entry.SetPlaceHolder("500kg")

	mass_set := widget.NewButton("Set", func() {fmt.Println("Mass has been set")})

	text_display_row := container.New(layout.NewHBoxLayout(), mass_label, mass_data)

	update_row := container.New(layout.NewGridLayout(2), mass_entry, mass_set)

	mass_group := container.New(layout.NewVBoxLayout(), text_display_row, update_row)

	return mass_group

}

func generate_side_menu(str binding.String) fyne.CanvasObject { 

	// First create buttons 

	group_1 := mass_button_group(str)
	group_2 := mass_button_group(str)

	restart_sim_button := widget.NewButton("Restart Simulation", func() {fmt.Println("Restart Clicked")})
	start_button := widget.NewButton("Start!", func() {fmt.Println("Start sim clicked")})

	menu := container.New(
		layout.NewVBoxLayout(),
		group_1,
		group_2, 
		layout.NewSpacer(), 
		restart_sim_button, 
		start_button)

	return menu
}


func main() { 

	// First start by creating an instance
	// of the app itself
	a := app.New() 

	// Then create the main window
	main_window := a.NewWindow("Gomentum")

	main_window.SetMaster() // Now it's the master window

	// Let's also resize it for
	// default behavior 
	main_window.Resize(fyne.NewSize(1000, 500))


	// text1 := canvas.NewText("Helo", color.White) 
	// text2 := canvas.NewText("There", color.White) 
	rect := canvas.NewRectangle(color.RGBA{R: 255, G:0, B: 0, A:255})


	rect.SetMinSize(fyne.NewSize(1200, 400))

	str := binding.NewString()
	go func() {
		dots := "....."
		for i := 5; i >= 0; i-- {
			str.Set("Count down" + dots[:i])
			time.Sleep(time.Second)
		}
		str.Set("Blast off!")
	}()

	side_menu := generate_side_menu(str)

	content := container.New(layout.NewHBoxLayout(), side_menu, rect) 

	// text4 := canvas.NewText("centered", color.White) 

	// centered:= container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())

	main_window.SetContent(content)

	main_window.Show() 

	a.Run() 

}