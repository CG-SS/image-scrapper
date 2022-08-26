package img_formats

type ImageFormat int64

const (
	JPEG ImageFormat = iota
	JPG
	GIF
	TIFF
	PSD
	AI
	INDD
	RAW
	PNG
	SVG
	UNKNOWN
)

func AllFormatsString() []string {
	var fmtString []string

	for _, format := range AllFormats() {
		fmtString = append(fmtString, format.String())
	}

	return fmtString
}

func AllFormats() []ImageFormat {
	return []ImageFormat{
		JPEG,
		JPG,
		GIF,
		TIFF,
		PSD,
		AI,
		INDD,
		RAW,
		SVG,
		PNG,
	}
}

func (i ImageFormat) String() string {
	switch i {
	case JPEG:
		return "JPEG"
	case JPG:
		return "JPG"
	case GIF:
		return "GIF"
	case TIFF:
		return "TIFF"
	case PSD:
		return "PSD"
	case AI:
		return "AI"
	case INDD:
		return "INDD"
	case RAW:
		return "RAW"
	case SVG:
		return "SVG"
	case PNG:
		return "PNG"
	}
	return "Unknown"
}

func ParseImageFormat(s string) ImageFormat {
	switch s {
	case "JPEG":
		return JPEG
	case "JPG":
		return JPG
	case "GIF":
		return GIF
	case "TIFF":
		return TIFF
	case "PSD":
		return PSD
	case "AI":
		return AI
	case "INDD":
		return INDD
	case "RAW":
		return RAW
	case "SVG":
		return SVG
	case "PNG":
		return PNG
	}
	return UNKNOWN
}
