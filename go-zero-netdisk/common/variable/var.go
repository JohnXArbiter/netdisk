package variable

import (
	"lc/netdisk/common/constant"
	"regexp"
)

var (
	Upattern, _ = regexp.Compile("^[a-zA-Z0-9]{6,20}$")
	Ppattern, _ = regexp.Compile("^[a-zA-Z0-9!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]{6,20}$")
)

var (
	DocSuffix = map[string]bool{
		".doc":  true,
		".docx": true,
		".xlsx": true,
		".xls":  true,
		".ppt":  true,
		".pptx": true,
		".pdf":  true,
		".txt":  true,
	}

	ImageSuffix = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".tif":  true,
		".tiff": true,
		".webp": true,
	}

	AudioSuffix = map[string]bool{
		".mp3":  true,
		".wav":  true,
		".aac":  true,
		".flac": true,
		".ogg":  true,
		".wma":  true,
		".m4a":  true,
		".amr":  true,
		".ape":  true,
	}

	VideoSuffix = map[string]bool{
		".avi":  true,
		".rm":   true,
		".rmvb": true,
		".mpeg": true,
		".mpg":  true,
		".mp4":  true,
		".wmv":  true,
		".mov":  true,
		".flv":  true,
		".mkv":  true,
		".3gp":  true,
		".m4v":  true,
	}
)

func GetTypeByBruteForce(ext string) int8 {
	if DocSuffix[ext] {
		return constant.TypeDocs
	} else if ImageSuffix[ext] {
		return constant.TypeImage
	} else if AudioSuffix[ext] {
		return constant.TypeAudio
	} else if VideoSuffix[ext] {
		return constant.TypeVideo
	} else {
		return constant.TypeOther
	}
}

var ShareExpireType = map[int8]int64{
	0: 0,
	1: 86400,
	2: 604800,
	3: 2592000,
}
