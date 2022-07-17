package main
import (
	"fmt"
	"skybot/ducksearch"
	"net/url"
)

func main() {
	w, err := wikimedia.New("http://en.wikipedia.org/w/api.php")
if err != nil {
	fmt.Println(err)
}
f := url.Values{
	"action": {"query"},
	"prop":   {"extracts"},
	"titles": {"kerala|Procrastination"},
}
res, err := w.Query(f)
if err != nil {
	 fmt.Println(err)
}

fmt.Println(res.Query.Pages)
// for _, v := range res.Query {
// 	fmt.Println(v.Title, "-", v.Extract)
// }

}
