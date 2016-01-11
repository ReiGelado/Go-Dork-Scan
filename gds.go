package main

import (
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"net/http"
	"net/url"
	"regexp"
	"flag"
	"strconv"
)

var dork_comando string
var paginas_comando int
var update_comando *bool
var buscador_comando string
var versao_script float64
var versao_script_string string
var regex_bing string
var regex_google string
var regex_update string

func erro(e error){
	if e != nil{
	    panic(e)
	}
}
func banner(){
	fmt.Println("[+]Go Dork Scan v" + versao_script_string + " by ReiGel_ado")
	fmt.Println("[+]Buscadores Disponiveis:\n[+]Bing\n[+]Google")
}
func argumentos_variaveis(){
	versao_script = 0.3
	versao_script_string = "0.3"
	flag.StringVar(&dork_comando,"dork","noticia.php?id=1"," - Sua Dork!")
	flag.StringVar(&buscador_comando,"buscador","bing"," - Buscador para puxar as dorks!")
	flag.IntVar(&paginas_comando,"paginas",1," - Numero de paginas para o script acessar!")
	update_comando = flag.Bool("update",false," - Verifica atualizações do script!")
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
func parser(html string,regex_entrada string)[][]string{
	regex := regexp.MustCompile(regex_entrada)
	resultado := regex.FindAllStringSubmatch(html,-1)[0:]
	return resultado
}
func update(versao string){
	regex_update = "<br><b><i>Versão do Script : (.*?) <br></i></b></p>"
	fmt.Println("[+]Versão do script(local) : ",versao)
	html_git := html_download("https://github.com/ReiGelado/Go-Dork-Scan/blob/master/README.md")
	resultado_versao := parser(html_git,regex_update)
	fmt.Println("[+]Versão do git           : ",resultado_versao[0][1])
	versao_git_float, err := strconv.ParseFloat(resultado_versao[0][1], 64)
	erro(err)
	if versao_script > versao_git_float{
		fmt.Println("[+]Versão Atualizada!")
	}else if versao_script == versao_git_float{
		fmt.Println("[+}Versão igual a do git!")
	}else {
		arquivo,err := os.Create("master.zip")
		erro(err)
		defer arquivo.Close()
		arquivo_down , err := http.Get("https://github.com/ReiGelado/Go-Dork-Scan/archive/master.zip")
		erro(err)
		defer arquivo_down.Body.Close()
		escreve_arquivo , err := io.Copy(arquivo,arquivo_down.Body)
		erro(err)
		_ = escreve_arquivo
		fmt.Println("[+]Download da nova atualização completo!")
	}
}
func bing(paginas int) []string {
	var resultado_slice_1 []string
	regex_bing = "</li><li class=\"b_algo\"><h2><a href=\"(.*?)\" h=\"ID=SERP,"
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <= 1 {
		recebe_download := html_download("http://www.bing.com/search?q=" + dork_escaped)
		resultado := parser(recebe_download,regex_bing)
		for i := 0; i < 10; i++ {
			resultado_slice_1 = append(resultado_slice_1,resultado[i][1])
		}
	} else if paginas > 1 {
		for pa := 1; pa <= paginas; pa++ {
			pa_str := strconv.Itoa(pa) //converta ela de int para str
			url_paginas := ("http://www.bing.com/search?q=" + dork_escaped + "&first=" + pa_str + "1" ) //url pa bing
			recebe_download := html_download(url_paginas)
			resultado := parser(recebe_download,regex_bing)
			for i := 0; i < 10; i++ {
				resultado_slice_1 = append(resultado_slice_1,resultado[i][1])	
			}
		}
	}
	return resultado_slice_1
}
func google(paginas int) []string {
	var resultado_slice_2 []string
	regex_google = `"><a href="/url\?q=(.*?)&amp;sa=U&amp;`
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <=1 {
		recebe_download := html_download("https://www.google.com.br/search?q=" + dork_escaped)
		resultado := parser(recebe_download,regex_google)
		for i := 0 ; i < 10; i++{
			url_unescaped,err := url.QueryUnescape(resultado[i][1])
			erro(err)
			resultado_slice_2 = append(resultado_slice_2,url_unescaped)
		}
	}else if paginas > 1{
		for pa := 1;pa <= paginas;pa++{
			pa_str := strconv.Itoa(pa)
			url_paginas := ("https://www.google.com.br/search?q=" + dork_escaped + "&start=" + pa_str + "0")//ulr pa google
			recebe_download := html_download(url_paginas)
			resultado := parser(recebe_download,regex_google)
			for i := 0; i < 9;i++{
				url_unescaped,err := url.QueryUnescape(resultado[i][1])
				erro(err)
				resultado_slice_2 = append(resultado_slice_2,url_unescaped)
			}
		}
	}
	return resultado_slice_2
}
func main() {
	argumentos_variaveis()
	banner()
	if *update_comando == true{
		update(versao_script_string)
		os.Exit(0)
	}
	if buscador_comando == "google"{
		result := google(paginas_comando)
		fmt.Println("[+]Resultado:")
		for i := 0;i < len(result);i++{
			fmt.Println("[+]Link:",result[i])
		}
	}else if buscador_comando == "bing"{
		result := bing(paginas_comando)
		fmt.Println("[+]Resultado:")
		for i := 0;i < len(result);i++{
			fmt.Println("[+]Link:",result[i])
	}
}
}
