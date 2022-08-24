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
	UNKNOWN
)

func AllFormatsString() []string {
	return []string{
		JPEG.String(),
		JPG.String(),
		GIF.String(),
		TIFF.String(),
		PSD.String(),
		AI.String(),
		INDD.String(),
		RAW.String(),
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
	}
	return UNKNOWN
}
