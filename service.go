package internal_proxy

import (
	"fmt"
	"github.com/phayes/freeport"
	"io/ioutil"
	"net/http"
)

// Serve will start and serve the process, it also creates the .htaccess file
func Serve() error {
	port, err := freeport.GetFreePort()
	if err != nil {
		return err
	}
	err = createFile(port)
	if err != nil {
		return err
	}
	fmt.Printf("listen on %d\n",port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return nil
}

// createFile will create the .htaccess file which redirects all traffic to the service
func createFile(port int) error {
	file := fmt.Sprintf("RewriteEngine on\nRewriteRule  (.*)  http://localhost:%d/$1  [P,L]\n", port)
	err := ioutil.WriteFile(".htaccess", []byte(file), 0644)
	if err != nil {
		return err
	}
	return nil
}
