package files

import (
	"os"

	"alukart32.com/tools"
)

func Copy(dstFile, srcFile string) {
	src, e := os.Open(srcFile)
	tools.CheckErr(e)
	defer src.Close()

	dst, e := os.OpenFile(dstFile, os.O_APPEND|os.O_WRONLY, 0644)
	tools.CheckErr(e)
	defer dst.Close()

	dst.ReadFrom(src)
}
