package menu

import "github.com/louisevanderlith/droxolite/menu"

func FullMenu() *menu.Menu {
	m := menu.NewMenu()
	m.AddItem(menu.NewItem("home", "/", "Home", nil))

	return m
}
