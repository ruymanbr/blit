<!-- ⚠️ This README has been generated from the file(s) "blueprint.md" ⚠️-->Welcome to @blit/readme. This is version 0.0.2!

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
	<li><a href="#4-production-readiness">4. Frontend demonstration (React App UI + API)</a></li>
</ul>
</div>
<h2>1-Installation</h2>
<h3>Requirements</h3>
<ul>
	<li>Ubuntu 16.04 or newer (tested in Ubuntu 16.04)</li>
	<li>Go version 1.16 or newer <sub>(Go to: <a href="https://golang.org/doc/install">Go install</a> to learn how to install it)</sub></li>
	<li>Git installed in your system <sub>(To see instructions on how to install Git in Ubuntu go to: <a href="https://github.com/git-guides/install-git#install-git-on-linux">Install Git on Linux</a></sub></li>
</ul>
<br />
<h3>Steps</h3>
<p>Open a Console in Ubuntu and follow this steps</p>
<br />
<ul>
	<li>cd ~</li>
	<li>mkdir temp && cd temp</li>
	<li>git clone https://github.com/ruymanbr/blit.git</li>
	<li>cd blit</li>
	<li>go build -o blit</li>	
</ul>
<h4>CLI</h4>
<ul>
	<li>i.e.: go run main.go /path/to/folder/</li>
	<sub> -- or ./blit /path/to/folder/</sub>
</ul>
<br />
<p>You should see a list of files and folders from the cloned git repo in your console, ordered by size (descending size)</p>
<br />
<h4>Frontend from CLI</h4>
<ul>
	<li>go run main.go</li>
	<sub> -- or ./blit</sub>
</ul>	
<br />
<p>You should see how your default browser (only Linux) is opened and the Frontend should be visible to operate on.</p>
<br />
<h2>2-Getting Started</h2>
<h3>Testing the package</h3>

<p>From CONSOLE, once inside BLIT directory, run the following command</p>
<ol>
	<li>go test</li> 
	<sub>(or 'go test blit_test.go')</sub>
</ol>
<br />
<h3>Go documentation. GODOC</h3>
<p>Documentation has been provided in this repository in DOC, PDF and godoc format. If you don't have GODOC installed in your system and want to use it to read the program documentation instead PDF/DOC fileformats, follow the next instructions.</p>
<br />
<p>Inside BLIT root directory run the following commands:</p>
<ol>
	<li>go get golang.org/x/tools/cmd/godoc</li>
	<li>cd pkg/blit_cli</li>
	<li>go doc -all blit_cli.go</li>
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
<h3>Discussion</h3>
<p>Program can be used through CLI or implementating it through its API into a frontend (tested on REACT-app frontend)</p>
<br />
<p>You can execute the demonstration in the next section</p>
<br />
<h3>Recommendations for backend</h3>
<ul><li>Stability and performance improvements</li>
		<ul>
			<li>Implementation of goroutines for parallel/concurrent executions of backend instances (More requests whenever more threads can run concurrently)</li>
			<li>Implementation of channeling and buffering in channels to further expand stability and capacity of backend</li>
			<li>Separation of blit_frontend program to run independently of backend server</li>
			<li>Implementation of parallelization/concurrency/process waiting lists (workers), in frontend to increment the capacity to deliver connections to and from backend to more clients at the same time</li>
			<li>Implementation of workers for a PWA version to work and save state even on loss of connection, retrying communication/processes as soon as connection is available</li>
			<li>Implementation of methods and functions that have ability to recover or retry the operation that failed through API and/or CLI</li>
			<li>Further testing should be implemented: FastSwitchSli function have no test implemented yet</li>
			<li>Further Benchmarking tests should be implemented. This should be done to proper evaluate performance. Specially for scaling up and remote using this program.</li>
			<sub>(See <a href="https://golang.org/pkg/testing/">GO Testing</a>)</sub>
			<li>Some more error handling in code could be added, specially with the use of Recover for increased stability</li>
			<sub>It's fairly managed, but some user entering parameters in some cases maight develop in error during runtime. Hence the need for more testing</sub>
			<li>Human error workarounds</li>
			<sub>The ability to work even when some minor mistakes are entered when calling the program (via API or CLI). Detection of different quotation symbols, lower/upper casing letters in the middle, missing some symbols etc</sub>
			<li>Double check and review that code closes every open file or path to improve memory management</li>
		</ul>
	<li>Complexity and Coverage</li>
		<ul>
			<li>More argument complexity in BLIT CLI (Adding flag capacity with flag package)</li>
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
		</ul>		
	<sub>Full example could be: blit -path="/usr/local/go" -size="asc" -filter="*.go"</sub>
	<li>Documentation:</li>
		<ul>
			<li>Examples should be added in blit_test.go to be displayed in Blit documentation (godoc)</li>
		</ul>
	<li>Simplification of functions:</li>
		<ul>
			<li>EncapData() function from blit_cli package is too large. E.g.: Too many parameters returned.</li>
			<li>Code could be fragmented even more to specialize it to make this more modular and reusable. Although the rest of the code it's fairly optimized</li>
			<li>Modularization</li>
			<sub>Subdividing code for different variations of parameters being passed into the program could lead to a faster minor programs that could be run separately. That would lead to a better integration with other programs (API calls, etc)</sub>
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
	<li>Logging ability</li>
	<sub>When errors happen, the ability to log and communicate, even through cloud APIs would be nice to have. Third party implementations with this program could get feedback and also recover from error handling calls to this aplication (maybe calling a 2nd option if this is unavailable due to some error, or format taken isn't parametrized according to user/app call, but the 2nd option is)</sub>
	<li>Go legacy compatibility</li>
	<sub>Code could be implemented to work with older versions. Proper error handling for this cases, and further testing should be implemented to make sure it has backwards compatibility</sub>	
	<li>Security:</li>
	<ul>
		<li>i.e: Handling calls when system where program is installed but user calling (API) isn't allowed to see folder content</li>
		<sub>Better encapsulation and isolation of data</sub>
		<li>Implementation of Secrets (pub key, priv key, ...) to exchange information between frontend and backend. Encryption, etc</li>
		<li>Database implementation for backend server</li>
		<ul>
			<li>Some packages are in the backend to use SQL for checking and serving data from the backend if necessary (i.e: User auth, checking credentials, etc)</li>
		</ul>
	</ul>
</ul>
<h2>4-Frontend demonstration</h2>
<sub>(React App UI + API)</sub>

