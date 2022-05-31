package termtitle

func SetTitle(title string) error {
	return setTitle(title)
}

func GetTitle() (string, error) {
	return getTitle()
}
