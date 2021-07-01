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
	<li>go test</li> 
	<sub>(or 'go test blit_test.go')</sub>
</ol>
<br />
<h3>Go documentation. GODOC</h3>
<br />
<p>Documentation has been provided in this repository in DOC, PDF and godoc format. If you don't have GODOC installed in your system and want to use it to read the program documentation instead PDF/DOC fileformats, follow the next instructions.</p>
<br />
<p>Inside BLIT root directory run the following commands:</p>
<ol>
	<li>go get golang.org/x/tools/cmd/godoc</li>
	<li>cd pkg</li>
	<li>go doc -all blit.go</li>
	<sub>Displays in console</sub>
</ol>
<br />
<p>You can also see it in your browser using 'godoc' command as follows</p>
<ol>
	<li>export BLIT_PATH=/here/path/to/blit/</li>	
	<li>godoc -http=:8080 -goroot=$BLIT_PATH && x-www-browser http://localhost:8080</li>
</ol>
<br />
<h2>3-Production Readiness</h2>
<br />
<h3>Discussion</h3>
<br />
<p>Program can be used through CLI or implementating it through its API into a frontend (tested on REACT-app from Redux)</p>
<br />
<p>You can execute the demonstration in the next section</p>
<br />
<h3>Recommendations for backend</h3>
<br />
<ul>
	<li>Further testing could be implemented.</li>
	<sub>(See <a href="https://golang.org/pkg/testing/">GO Testing</a>)</sub>
	<li>Some more error handling in code could be added</li>
	<sub>It's fairly managed, but some user enter parameters could fire up some errors that could develop in function called inside the program.</sub>
	<li>Stability improvements</li>
		<ul>
			<li>Implementation of methods that have ability to recover or retry the operation that failed</li>
		</ul>
	<li>More argument complexity in BLIT CLI.</li>
	<sub>(Parameters for different options)</sub>
		<sub>Examples:</sub>
		<ul>			
			<li>blit -path="/path/to/a/folder/"</li>
			<sub>Instead of taking the first parameter as it goes now (i.e.: blit /path/here)</sub>
			<li>blit -size="asc"</li>
			<sub>To display files from smaller to bigger in size</sub>
			<li>blit -date="asc"</li>
			<sub>To display files from more recently modified to last date modified</sub>
			<li>blit -date="desc"</li>
			<sub>The opposite</sub>
			<li>blit -filter="abcd"</li>
			<sub>To display only files with 'abcd' in their names (0000abcd.doc; 98abcd_this_too.pdf) ... and so on</sub>
			<li></li>
		</ul>
		<sub>Full example could be: blit -path="/usr/local/go" -size="asc" -filter="*.go"</sub>
	<li>Simplification of functions:</li>
		<ul>
			<li>Several functions are in a beta state. E.g.: Too many parameters returned.</li>
			<li>Code could be fragmented even more to specialize some functions to make this more modular and reusable. Although it's fairly optimized</li>
		</ul>
	<li>Further Beautification:</li>
		<ul>
			<li>Colouring Folders instead of column specifying whenever a "file" is in reality a folder</li>
			<li>Colouring files in different tones depending on file size</li>
			<sub>(biggest in different tones of red, lighter in different tones of green, etc)</sub>
		</ul>
	<li>OS compatibility</li>
		<ul>
			<li>Make paths to be compatible with main OS in market (Windows, Mac, ...)</li>
			<li>Backwards compatibility with older versions of OS would be nice to have</li>
		</ul>
	<li>Loggin ability</li>
	<sub>When errors happen, the ability to log and communicate, even through cloud APIs would be nice to have. Third party implementations with this program could get feedback and also recover from error handling calls to this aplication (maybe calling a 2nd option if this is unavailable due to some error, or format taken isn't parametrized according to user/app call, but the 2nd option is)</sub>
	<li>Go older version's Compatibility</li>
	<sub>Code could be implemented to work with older versions. Proper error handling for this cases, and further testing should be implemented to make sure it has backwards compatibility</sub>
	<li>User error workarounds</li>
	<sub>The ability to work even when some minor mistakes are entered when calling the program (via API or CLI). Detection of different quotation symbols, lower/upper casing letters in the middle, missing some symbols etc</sub>
	<li>Modularization</li>
	<sub>Subdividing code for different variations of parameters being passed into the program could lead to a faster minor programs that could be run separately. That would lead to a better integration with other programs (API calls, etc)</sub>
	<li>Security:</li>
		<ul>
			<li>i.e: Handling calls when system where program is installed but user calling (API) isn't allowed to see folder content</li>
			<li></li>
		</ul>
</ul>
<br />
<h3>Recommendations for Frontend</h3>
<ul>
	<li></li>
		<ul>
			<li></li>
			<li></li>
			<li></li>
			<li></li>
		</ul>
	<li></li>
</ul>
{{ load:readme/4-frontend-demonstration.md }}

