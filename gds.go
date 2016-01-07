package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"net/url"
	"regexp"
	"flag"
)

var dork_comando string

func erro(e error){
	if e != nil{
	    panic(e)
	}
}
func banner(){
	fmt.Println("[+]Go Dork Scan Beta v0.1 by ReiGel_ado")
	fmt.Println("[+]Buscadores Disponiveis:\n[+]Bing\n[+]Google(Futuramente)\n")
	flag.StringVar(&dork_comando,"dork","noticia.php?id=1"," - Sua Dork")
	flag.Parse() //Carega os argumentos
}
func escreve(valor string,arquivo string){ //implementação futura,não ligue pro arquivo criado...
	wa , err := os.Create(arquivo)
	erro(err)
	defer wa.Close()
	wa.WriteString(valor)
	wa.Sync()
}
func html_download(url_site string,dork string) string {
	dork_escaped := url.QueryEscape(dork)
	html_down,erro_down := http.Get(url_site + dork_escaped)
	erro(erro_down)
	defer html_down.Body.Close()
	html_body,body_erro := ioutil.ReadAll(html_down.Body)
	erro(body_erro)
	string_body := string(html_body)
	return string_body
}
func main() {
	banner()
	recebe_download := html_download("http://www.bing.com/search?q=",dork_comando)
	escreve(recebe_download,"html.txt")
	//fmt.Println(recebe_download)
	regex := regexp.MustCompile("</li><li class=\"b_algo\"><h2><a href=\"(.*?)\" h=\"ID=SERP,")
	result := regex.FindAllStringSubmatch(recebe_download,-1)[0:]
	for i := 1; i <= 9; i++ { //Retorna [(completo,url)]
		fmt.Println(result[i][1])
	}
}
