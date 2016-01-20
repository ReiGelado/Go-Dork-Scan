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
var saida_comando string
var proxy_comando string
var useragent_comando string
var versao_script float64
var versao_script_string string
var result []string
var resultado_slice_1 []string //Vou otimizar isso depois kkkk
var resultado_slice_2 []string
var resultado_slice_3 []string
var resultado_slice_4 []string
var resultado_slice_5 []string
var regex_bing string
var regex_google string
var regex_update string
var regex_duckduckgo string
var regex_yahoo string
var regex_ask string

func erro(e error){
	if e != nil{
	    panic(e)
	}
}
func banner(){
	fmt.Println("[+]Go Dork Scan v" + versao_script_string + " by ReiGel_ado")
	fmt.Println("\n[+]Buscadores Disponiveis:\n[+]Bing\n[+]Google\n[+]DuckDuckGo(Sem Paginas)\n[+]Yahoo\n[+]Ask\n")
}
func argumentos_variaveis(){
	flag.StringVar(&dork_comando,"dork","noticia.php?id=1"," - Sua Dork!")
	flag.StringVar(&buscador_comando,"buscador","bing"," - Buscador para puxar as dorks!")
	flag.StringVar(&saida_comando,"saida","false"," - Um saida para as dorks!")
	flag.StringVar(&useragent_comando,"user-agent","false"," - Modifica o user-agent do script!")
	//flag.StringVar(&proxy_comando,"proxy","false"," - Seta um proxy no script(BETA)!")
	flag.IntVar(&paginas_comando,"paginas",1," - Numero de paginas para o script acessar!")
	update_comando = flag.Bool("update",false," - Verifica atualizações do script!")
	flag.Parse() //Carega os argumentos
	versao_script = 0.5
	versao_script_string = "0.5"
	if useragent_comando == "false"{
		useragent_comando = "Go-http-client/1.1"
	}else {
		fmt.Println("[+]User-Agent : " + useragent_comando)
		fmt.Println("[+]Lembrando que o User-Agent influencia no resultado das DORKS!\n")
	}
}
func escreve_slice(nome_do_arquivo string,array_slice []string){
	arquivo_w,err := os.Create(nome_do_arquivo)
	erro(err)
	defer arquivo_w.Close()
	for i := range array_slice{
		arquivo_w.WriteString(array_slice[i] + "\n")
	}
	arquivo_w.Sync()
	if _, err := os.Stat(nome_do_arquivo); err == nil { 
		fmt.Println("\n[+]Dork Salvas em :" + nome_do_arquivo)
	} else if err != nil {
		fmt.Println("[+]Ocorreu um erro na escrita do arquivo....")
		erro(err)
	}
}
func html_download(url_site string) string {
	navegador := &http.Client{}
	cabecalho , err := http.NewRequest("GET",url_site, nil)
	erro(err)
	cabecalho.Header.Set("User-Agent",useragent_comando)
	html_down,err := navegador.Do(cabecalho)
	erro(err)
	defer html_down.Body.Close()
	html_body,err := ioutil.ReadAll(html_down.Body)
	erro(err)
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
		fmt.Println("[+]Verifique o arquivo master.zip !")
	}
}
func bing(paginas int) []string {
	regex_bing = "</li><li class=\"b_algo\"><h2><a href=\"(.*?)\" h=\"ID=SERP,"
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <= 1 {
		recebe_download := html_download("http://www.bing.com/search?q=" + dork_escaped)
		resultado := parser(recebe_download,regex_bing)
		for i := range resultado{
			resultado_slice_1 = append(resultado_slice_1,resultado[i][1])
		}
	} else if paginas > 1 {
		for pa := 1; pa <= paginas; pa++ {
			pa_str := strconv.Itoa(pa) //converta ela de int para str
			url_paginas := ("http://www.bing.com/search?q=" + dork_escaped + "&first=" + pa_str + "1" ) //url pa bing
			recebe_download := html_download(url_paginas)
			resultado := parser(recebe_download,regex_bing)
			for i := range resultado{
				resultado_slice_1 = append(resultado_slice_1,resultado[i][1])	
			}
		}
	}
	return resultado_slice_1
}
func google(paginas int) []string {
	fmt.Println("[+]Buscador Google em função beta...")
	fmt.Println("[+]CAPTCHA aparece de vez em quando...")
	regex_google = `"><a href="/url\?q=(.*?)&amp;sa=U&amp;`
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <=1 {
		recebe_download := html_download("https://www.google.com.br/search?q=" + dork_escaped + "&oq="+ dork_escaped + "&gws_rd=cr,ssl&client=ubuntu&ie=UTF-8")
		resultado := parser(recebe_download,regex_google)
		for i := range resultado{
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
			for i := range resultado{
				url_unescaped,err := url.QueryUnescape(resultado[i][1])
				erro(err)
				resultado_slice_2 = append(resultado_slice_2,url_unescaped)
			}
		}
	}
	return resultado_slice_2
}
func meu_pato(paginas int)[]string{
	regex_duckduckgo = `<a rel=\"nofollow\" href=\"(.*?)\">`
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <= 1 {
		recebe_download := html_download("https://duckduckgo.com/html/?q=" + dork_escaped)
		resultado := parser(recebe_download,regex_duckduckgo)
		for i := range resultado{
			url_unescaped,err := url.QueryUnescape(resultado[i][1])
			erro(err)
			resultado_slice_3 = append(resultado_slice_3,url_unescaped)
		}
	}else if paginas >1 {
		fmt.Println("[+]Pensando em como implementar paginas no Duck 0-0!")
		return nil
	}
	return resultado_slice_3
}
func yahoo(paginas int)[]string{
	regex_yahoo = `\" ac-algo ac-21th lh-15\" href=\"(.*?)\" target=\"_blank`
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <= 1{
		recebe_download := html_download("https://search.yahoo.com/search?p=" + dork_escaped)
		resultado := parser(recebe_download,regex_yahoo)
		for i := range resultado{
			resultado_slice_4 = append(resultado_slice_4,resultado[i][1])
		}
	}else if paginas > 1{
		for pa := 1;pa <= paginas;pa++{
			pa_str := strconv.Itoa(pa)
			url_paginas := ("https://search.yahoo.com/search?p=" + dork_escaped + "&ei=UTF-8&b=" + pa_str + "1")
			recebe_download := html_download(url_paginas)
			resultado := parser(recebe_download,regex_yahoo)
			for i := range resultado{
				resultado_slice_4 = append(resultado_slice_4,resultado[i][1])
			}
		}
	}
	return resultado_slice_4
}
func ask(paginas int)[]string{
	regex_ask = `<a class=\"web-result-title-link\" href=\"(.*?)\" onmousedown=\"uaction\(this`
	dork_escaped := url.QueryEscape(dork_comando)
	if paginas <= 1{
		recebe_download := html_download("http://www.ask.com/web?q=" + dork_escaped)
		resultado := parser(recebe_download,regex_ask)
		for i := range resultado{
			resultado_slice_5 = append(resultado_slice_5,resultado[i][1])
		}
	}else if paginas > 1{
		for pa := 1;pa <= paginas;pa++{
			pa_str := strconv.Itoa(pa)
			url_paginas := ("http://www.ask.com/web?q=" + dork_escaped + "&page=" + pa_str)
			recebe_download := html_download(url_paginas)
			resultado := parser(recebe_download,regex_ask)
			for i := range resultado{
				resultado_slice_5 = append(resultado_slice_5,resultado[i][1])
			}
		}
	}
	return resultado_slice_5
}
func main() {
	banner()
	argumentos_variaveis()
	if *update_comando == true{
		update(versao_script_string)
		os.Exit(0)
	}
	if buscador_comando == "google"{
		result = google(paginas_comando)
	}else if buscador_comando == "bing"{
		result = bing(paginas_comando)
	}else if buscador_comando == "duck"{
		result = meu_pato(paginas_comando)
	}else if buscador_comando == "yahoo"{
		result = yahoo(paginas_comando)
	}else if buscador_comando == "ask"{
		result = ask(paginas_comando)
	}
	for i := range result{
		fmt.Println("[+]Link:",result[i])
	}
	if saida_comando != "false"{//poise....
		escreve_slice(saida_comando,result)
	}
}
