package create

import (
	"dao/wordpress/theme"
)

func Make(item string) {

		switch item {
			case "theme":
				theme.Scaffold()
		}
	}