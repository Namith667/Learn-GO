package main
//implementation of goroutines and channels
import (
	"fmt"
	"net/http"
	"time"
)
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)

	for _, link:=  range links{
		go checkLink(link,  c)
	}
	
	
	//channel listens in a looping manner till size limit of slice in checkLink
	//receive channel 
	
	//makes use of closures in anonymous function
	for l := range c {
		go func (link string){
			time.Sleep(5*time.Second)
			checkLink(link, c)
		}(l)// pass `l` explicitly to avoid sharing issues
	}
}

func checkLink(link string, c chan string){
	_,err := http.Get(link)
	if err != nil{
		fmt.Println(link, "might be down")
		c <- link//  we can pass message or link(argument)
		return
	}
	fmt.Println(link," is up")
	c <- link
}