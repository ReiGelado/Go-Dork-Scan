package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"net/url"
	"regexp"
	"flag"
	"strconv"
)

var dork_comando string
var paginas_comando int

func erro(e error){
	if e != nil{
	    panic(e)
	}
}
func banner(){
	fmt.Println("[+]Go Dork Scan v0.2 by ReiGel_ado")
	fmt.Println("[+]Buscadores Disponiveis:\n[+]Bing\n[+]Google(Futuramente)")
	flag.StringVar(&dork_comando,"dork","noticia.php?id=1"," - Sua Dork!")
	flag.IntVar(&paginas_comando,"paginas",1," - Numero de paginas para o script acessar!")
	flag.Parse() //Carega os argumentos
}
func escreve(valor string,arquivo string){ //implementação futura,não ligue pro arquivo criado...
	wa , err := os.Create(arquivo)
	erro(err)
	defer wa.Close()
	wa.WriteString(valor)
	wa.Sync()
}
func html_download(url_site string) string {
	html_down,erro_down := http.Get(url_site)
	erro(erro_down)
	defer html_down.Body.Close()
	html_body,body_erro := ioutil.ReadAll(html_down.Body)
	erro(body_erro)
	string_body := string(html_body)
	return string_body
}
func parser(html string,)[][]string{
	regex := regexp.MustCompile("</li><li class=\"b_algo\"><h2><a href=\"(.*?)\" h=\"ID=SERP,")
	resultado := regex.FindAllStringSubmatch(html,-1)[0:]
	return resultado
}
func bing(paginas int) []string {
	var nova_string []string
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <= 1 {
		recebe_download := html_download("http://www.bing.com/search?q=" + dork_escaped)
		resultado := parser(recebe_download)
		for i := 0; i < 10; i++ {
			nova_string = append(nova_string,resultado[i][1])
		}
	} else if paginas > 1 {
		for pa := 1; pa <= paginas; pa++ {
			pa_str := strconv.Itoa(pa) //converta ela de int para str
			url_paginas := ("http://www.bing.com/search?q=" + dork_escaped + "&first=" + pa_str + "1" ) //cria a url
			recebe_download := html_download(url_paginas)
			resultado := parser(recebe_download)
			for i := 0; i < 10; i++ {
				nova_string = append(nova_string,resultado[i][1])	
			}
		}
	}
	return nova_string
}
func main() {
	banner()
	result := bing(paginas_comando)
	fmt.Println("[+]Resultado:")
	for i := 1;i < len(result);i++{
		fmt.Println("[+]Link:",result[i])
	}
}
