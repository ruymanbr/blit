<!-- ⚠️ This README has been generated from the file(s) "blueprint.md" ⚠️-->Welcome to @blit/readme. This is version 0.0.1!

<p align="center">
  <img src="https://raw.githubusercontent.com/ruymanbr/blit/main/assets/blit_logo.png" alt="Logo" width="150" height="auto" />
</p>
<h1 align="center">@blit/readme</h1>
<p align="center">
		<a href="https://github.com/badges/shields"><img alt="Go version" src="https://img.shields.io/badge/Go-v1.16-blue" height="20"/></a>
<a href="https://github.com/badges/shields"><img alt="Platform" src="https://img.shields.io/badge/platform-linux-lightgrey" height="20"/></a>
	</p>

<p align="center">
  <b>List files and folders in a given path. Displays the next information: DIR y/n - Last modified date - file/folder name - size</b></br>
  <sub> Go here to see a demo <a href="https://raw.githubusercontent.com/ruymanbr/blit/main/assets/demo_blit.gif">https://raw.githubusercontent.com/ruymanbr/blit/main/assets/demo_blit.gif</a>.<sub>
</p>

<br />

<div id="toc_container" align="center">
<h2 class="toc_title" align="left">Content</h2>
<ul class="toc_list" align="left">
	<li><a href="#1-installation">1 Installation</a>  
	<li><a href="#2-getting-started">2. Getting Started</a></li>
	<li><a href="#3-production-readiness">3. Production Readiness</a></li>
</ul>
</div>
<h2>1-Installation</h2>
<br />
<h3>Requirements</h3>
<br />
<ul>
	<li>Ubuntu 16.04 or newer (tested in Ubuntu 16.04)</li>
	<li>Go version 1.16 or newer <sub>(Go to: <a href="https://golang.org/doc/install">Go install</a> to learn how to install it)</sub></li>
	<li>Git installed in your system <sub>(To see instructions on how to install Git in Ubuntu go to: <a href="https://github.com/git-guides/install-git#install-git-on-linux">Install Git on Linux</a></sub></li>
</ul>
<br />
<h3>Steps</h3>
<br />
<p>Open a Console in Ubuntu and follow this steps</p>
<br />
<ul>
	<li>cd ~</li>
	<li>mkdir temp && cd temp</li>
	<li>git clone https://github.com/ruymanbr/blit.git</li>
	<li>cd blit</li>
	<li>go build blit.go</li>
	<li>go run blit.go (or ./blit)</li>
</ul>
<br />
<p>You should see a lit of files and folders from the cloned git repo in your console</p>
<br />
<h2>2-Getting Started</h2>
<br />
<h3>Testing the package</h3>
<br />

<p>From CONSOLE, once inside BLIT directory, run the following command</p>
<ol>
	<li>go test (or go test blit_test.go)</li>
</ol>
<br />
<h3>Go documentation. GODOC</h3>
<br />
<p>Documentation has been provided in this repository in DOC, PDF and godoc format. If you don't have GODOC installed in your system and want to use it to read the program documentation instead PDF/DOC fileformats, follow the next instructions.</p>
<br />
<p>Inside Program BLIT, run the next commands:</p>
<ol>
	<li>go get golang.org/x/tools/cmd/godoc</li>
	<li>cd pkg</li>
	<li>go doc -all blit.go</li>
	<sub>Displays in console</sub>
</ol>
<br />
<p>You can also see it in your browser using 'godoc' command as follows</p>
	<li>export BLIT_PATH=/here/path/to/blit/</li>	
	<li>godoc -http=:8080 -goroot=$BLIT_PATH && x-www-browser http://localhost:8080</li>
</ol>
<br />


