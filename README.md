# Go Dork Scan Beta
Um script simples que tem a função de usar buscadores para retornar dorks.
<br><b><i>Versão do Script : 0.3 <br></b></i>
# Dependencias:
<b>Para usar o script você deve instalar a linguagem de programação golang,no ubuntu vocẽ pode instalar ela com<br>
o seguinte comando:</b>
<br><code>arthur@arthur-netbook:~$sudo apt-get install golang</code>
# Usage:
<b>Compilado:</b>
<br><code>arthur@arthur-netbook:~$./gds -dork=noticia.php?id=</code>
<br><b>Não compilado</b>(Ele vai compilar é rodar kk):
<br><code>arthur@arthur-netbook:~$go run gds.go -dork=noticia.php?id=</code>
# Compilando:
<b>Linux:</b>
<br><code>arthur@arthur-netbook:~$ go build gds.go</code>
<br><b>Windows:</b>
<br>Você pode tentar compilar para Windows(não portei e nem vou)com o seguinte comando:
<br><code>arthur@arthur-netbook:~$GOOS=windows GOARCH=386 go build -o gds.exe gds.go</code>
#Print
<b>Comando:</b>gds.go -dork=noticia.php?id=
<br><img src="http://i.imgur.com/o53ameH.png?1"></img>
<b>Comando:</b>>gds.go -dork=noticia.php?id= -paginas=2
<br><img src="http://i.imgur.com/UlhMyxf.png?1"></img>
<b>Comando:</b>gds.go -dork=noticia.php?id= -buscador=google -paginas=2
<br><i>Obs:No argymento -buscador vocẽ pode usar o "bing" tambem!</i>
<br><img src="http://i.imgur.com/ER9uv0g.png?1" ></img>
<b>Comando:</b>gds.go -update
<br><img src="http://i.imgur.com/uJNHYBc.png?1"></img>
