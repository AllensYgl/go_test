package main

import (
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type Handle func(write http.ResponseWriter, request *http.Request) error

//Treat error in the func and other func not notice the error only return error
// compel a func and return a func
func errWapper(handler Handle) func(write http.ResponseWriter, request *http.Request) { // the func execute in the return func
	return func(write http.ResponseWriter, request *http.Request) { //means use the func in the other func return value
		err := handler(write, request) // excute
		if err != nil {
			switch {
			case os.IsNotExist(err):
				http.Error(write, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			default:
				//http.Error(write, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}

// only run the func if face error only return error
func handlefile(write http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(write, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	write.Write(all)
	return nil
}

func main() {
	http.HandleFunc("/list/", errWapper(handlefile))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
