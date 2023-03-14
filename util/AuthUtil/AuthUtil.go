package AuthUtil

import "github.com/ronenzr/forreal.go/util/FileUtil"

var PublicKey = FileUtil.GetFileContent("./data/public.key")
