package charts

import (
	"embed"
	"io/fs"
)

var (
	//go:embed imagepullsecrets imagepullsecrets/templates/_helpers.tpl
	imagePullSecretsEmbed embed.FS

	// ImagePullSecrets exposes the imagepullsecrets chart using relative file paths from the chart root
	ImagePullSecrets fs.FS
)

func init() {
	var err error
	ImagePullSecrets, err = fs.Sub(imagePullSecretsEmbed, "imagepullsecrets")
	if err != nil {
		panic(err)
	}
}
