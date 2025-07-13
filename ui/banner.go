package ui

import "fmt"

const HelloBanner = `
╔════════════════════════════════════╗
║           MPV PLAYER CTRL          ║
║          	by MrPhoenix         	 ║
╚════════════════════════════════════╝

Привет! Что будем делать?

[1] Посмотреть список
[2] Добавить ссылку
[3] Запустить видео
[4] Запустить Cava
[5] Запустить EasyEffects
[6] Настройки
[0] Выход

> _

`

func ShowMainMenu() {
	fmt.Print(HelloBanner)
}
