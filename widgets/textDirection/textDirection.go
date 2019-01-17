package textDirection

type TextDirection string

const (
	//Ltr the text flows from left to right (e.g., English, French)
	Ltr TextDirection = "ltr"
	// Rtl the text flows from right to left (e.g. Arabic, Hebrew)
	Rtl TextDirection = "rtl"
)
