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
  <sub> Go here to see a demo <a href="https://raw.githubusercontent.com/ruymanbr/blit/main/assets/demos/blit_client_demo.gif">https://raw.githubusercontent.com/ruymanbr/blit/main/assets/demos/blit_client_demo.gif</a>.<sub>
</p>

<br />

<div id="toc_container" align="center">
<h2 class="toc_title" align="left">Content</h2>
<ul class="toc_list" align="left">
	<li><a href="#1-installation">1 Installation</a>  
	<li><a href="#2-getting-started">2. Getting Started</a></li>
	<li><a href="#3-testing">3. Testing</a></li>
	<li><a href="#4-documentation">4. Documentation</a></li>
	<li><a href="#5-production-readiness">5. Production Readiness</a></li>
</ul>
</div>
<h2>1-Installation</h2>
<h3>Requirements</h3>
<ul>
	<li>Ubuntu 16.04 or newer (tested in Ubuntu 16.04)</li>
	<li>curl (7.58.0 is latest stable)</li>
		<sub>To install it in UBUNTU use "apt-get install curl". For other linux systems check commands)</sub>
	<li>Go version 1.16 or newer <sub>(Go to: <a href="https://golang.org/doc/install">Go install</a> and follow the instructions)</sub></li>
	<li>IMPORTANT!! You should add "sudo" before the 2 instructions to unzip the file from GO INSTALL instructions</li>
	<li>IMPORTANT!! Also remember that you should adapt the command and/or go to the folder where you downloaded GO zip/tar file</li>
	<li>IMPORTANT!! You should add to your /home/path/.bashrc the "export ..." command for go BIN folder aswell. Else it'll be lost once your shell is closed</li>
	<li>Git installed in your system <sub>(To see instructions on how to install Git in Ubuntu go to: <a href="https://github.com/git-guides/install-git#install-git-on-linux">Install Git on Linux</a></sub></li>
	<li>Node.js (and npm >= 6.13.4) (that's the version it was tested on).</li>
		<sub>This will be necessary for the FRONTEND demo app I've made in React. To install it in UBUNTU follow the next instructions<a href="https://linuxize.com/post/how-to-install-node-js-on-ubuntu-18.04/">How to Install Node.js and npm on Ubuntu 18.04</a></sub>
	<li>Yarn >= 1.21.1 </li>
		
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
	<sub> -- or ./bin/blit /path/to/folder/</sub>
</ul>
<p>You should see a list of files and folders in your console, ordered by size (descending size)</p>
<br />
<h4>Blit Client installation</h4>
	<sub>(Development build)</sub>
<li>cd blit_frontend</li>
<h5>Installing general NPM dependancies</h5>
<sub>It'll install all dependencies found in package-lock.json (except individuals. Next line)</sub>
<sub>from blit_frontend directory run:</sub>
<li>npm ci</li>	
<h5>Installing local NPM dependancies. Run:</h5>
<li>npm install</li>
<h5>Installing Yarn (compatible with what's installed in blit_frontend. Run these commands:</h5>
<ul>
	<li>sudo apt remove cmdtest</li>
	<li>sudo apt remove yarn</li>
	<li>curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -</li>
	<li>echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list</li>
	<li>sudo apt-get update</li>
	<li>sudo apt-get install yarn -y</li>
	<li>yarn install</li>
</ul>
<p>Next sections will tell you how to test CLI, API and Client (foreground app)</p>

<h2>2-Getting Started</h2>
<h3>Testing CLI</h3>

<p>From CONSOLE, once inside BLIT directory, run the following commands</p>
<ol>
	<li>cd pkg/blit_cli</li>
	<li>go test -bench . -v</li>
	<sub>(or 'go test blit_cli_test.go -bench . -v')</sub>
</ol>
<h3>Testing Client (API + UI)</h3>
<li>Open 2 CONSOLE shells in Linux (i.e. Ubuntu)</li>
<li>Go to BLIT root folder in 1, and blit_frontend in 2nd. Run the following commands</li>
<ol>
	<li>Console 1: npm start server</li>
	<sub>Alternatively you can enter: "go run main.go"</sub>
	<sub> -- or ./bin/blit (if you are in BLIT root folder). There is a compiled version in BIN called "blit"</sub>
	<sub>IMPORTANT! If you have Google Chrome in Ubuntu, it comes with several bugs, like EXEC opening browser but not executing command. You'll need to follow this instructions: <a href="shorturl.at/sxzA1">How to xdg/open in Ubuntu Chrome</a></sub>
	<li>Console 2: yarn start</li>
	<sub>localhost:8080 and localhost:3000 (backend and client respectivel)</sub>
</ol>
<p>You can check in the CONSOLE 1 the request/response from CLIENT APP. </p>
<p>You can check http://localhost:8080/api/v1</p>
<p>You should expect a JSON with a message component saying "PONG" which verifies it's operating</p>
<p>Also, keep your console open and running. It'll show ROTER information like headers from responses through the endpoints. To check them from a frontend demo that is currently also available. See the next section.</p>
<p>CLIENT FOREGROUND APP: From localhost:3000 try entering PATHS in your current OS filesystem.</p>
<p> Happy Hacking!</p>



<h2>3-Testing</h2>

<p>Only unit tests have been provided for CLI. From CONSOLE, once inside BLIT directory, run the following commands</p>
<ol>
	<li>cd pkg/blit_cli</li>
	<li>go test -bench . -v</li>
	<sub>(or 'go test blit_cli_test.go -bench . -v')</sub>
</ol>
<p>Benchmarking tests at the end may vary depending on your computer. Lowest process has been tested around 30.000 ns/op. Fastest at 20 ns/op in one Laptop Intel i3, and desktop intel i7 with 8GB and 16GB, both SSD drives.</p>
<br />
<h2>4-Documentation</h2>
<p>Documentation has been provided in this repository in DOC, PDF and godoc HTML format. </p>
<br />
<p>Alternatively, inside BLIT root directory run the following commands:</p>
<ol>
	<li>go get golang.org/x/tools/cmd/godoc</li>
	<li>cd pkg/{pkg_name}</li>
	<li>go doc -all {pkg_name}.go</li>
	<sub>Displays in console</sub>
</ol>
<br />
<p>You can also see it in your browser using 'godoc' command (as follows)</p>
<ol>
	<li>export BLIT_PATH=/here/path/to/blit/</li>	
	<li>godoc -http=:6060 -goroot=$BLIT_PATH && x-www-browser http://localhost:6060</li>
	<sub>Remember that server + client use 8080 and 3000 respectively in this setup.</sub>
</ol>
<h3>API v1</h3>
<sub>(brief summary)</sub>
<p>API enpoints and protocols are very simplistic, due to the specifications submitted. It's a private API with 3 endpoints, which of whom uses only 1 to function:</p>
<p>Generally speaking, it's been made to accept GET, POST and OPTIONS, although it only works with POST, by receiving a JSON object. In this development phase, it's only taking first element as valid for "path" argument.</p>
<h4>Allowed (Protocols)</h4>
<li>Headers: "Content-Type", "application/json"</li>
<h4>Endpoints</h4>
<li>/api/v1 (GET)</li>
<p>i.e.:Checks for ONLINE/OFFLINE can be done</p>
<li>Methods:</li>
<ul>
	<li>GET</li>
	<p>Status Code: 200</p>
</ul>
<li>/api/v1/post</li>
<p>Entry point for POST/GET request</p>
<li>Methods:</li>
<ul>
	<li>GET</li>
	<p>Response:</p>
	<li>Status Code: 200</li>
	<li>POST</li>
	<ul>
		<li>Accepting: JSON object</li>
		<li>Parameters: </li>
		<p>JSON accepts 1 argument as 'key, value' inside JSON object: </p>
		<p>path: path_value    // of type STRING</p>
		<li>VALID OBJECT (JSON)? Then:</li>
		<ul>
			<li>VALID PATH submited? Then:</li>
			<ul>
				<li>VALID Request:</li>
				<ul>
					<p>Response:</p>
					<li>4 col array of type STRING with:</li>
					<li>key, value STRING</li>
				</ul>
				<p>Response format: {IsDir : value, LastM: value, FName: value, FSize_HR_Format: value}</p>
				<p>Status Code: 200</p>
				<li>INVALID Request: </li>
				<p>Response:</p>
				<ul>
					<li>{"message":"Unauthorized Path 401","error": err,}</li>
				</ul>
				<p>Notes: 'path' was processed but got an invalid answer from bg application, or an error</p>
			</ul>
			<li>NOT a VALID PATH submited? Then:</li>
			<li>INVALID Request: If not a valid "path" argument was submitted</li>
			<p>Response:</p>
			<ul>				
				<p>{"message":"Not found 404"}</p>
			</ul>
			<p>Notes: 'Path' was empty, or didn't got 'path' as key argyment in JSON</p>
		</ul>
		<li>INVALID Request (others): NOT VALID OBJECT // NOT a JSON object.</li>	
		<p>Response:</p>
		<ul>
			<p>{"message":"Internal Server Error 500","error": err,}</p>
		</ul>		
	</ul>
</ul>
<li>Other URIS or requests:</li>
<p>Response:</p>
<ul>				
	<p>{"message":"Not found 404"}</p>
</ul>
<br />
<h2>5-Production Readiness</h2>
<h3>Discussion</h3>
<p>Application is production ready (ish). Only with the specs provided. But I'd recommend further development. Which I comment in "What's next" section.</p>
<p>Blit v0.0.1 can be used through CLI or implemented it through its API into a frontend client app (tested on REACT-app for the DEMO)</p>
<br />
<p>You can execute the demonstration in the section <a href="#2-getting-started">'Getting Started'</a></p>
<br />
<h4>Steps to become PROD READY: (Minimum specs)</h4>
<li>React Client needs to abandon development status to be production ready. Right now is up for test, but isn't workable in a different machine or folder than the current one. It's a "demo" for foreground requests to the API bg system</li>
<li>BLIT Cli can be build and used as a CLI tool to list folders, but it's only fulfilling task minimum specs</li>
<li>Blit API is discussed below, and several changes should be made to become a REST API complete, when Blitadvance into a more broad spectrum of functionalities</li>
<h4>What's next? Broadening BLIT 0.0.1 to the next level and MANDATORY development to be PROD READY:</h4>

<h3>Some recommendations and improvements to achieve production Ready level</h3>
<ul>
	<li>CI/CD</li>
		<ul>Implementation of the API to a CI/CD friendly environment (Swagger, Swaggo?, Postman, ...) so it could integrate with CI/CD easily.</ul>
		<li>Documentation for the API should be provided (Prod Ready. CI/CD purposes, etc)</li>
		<li>More and better testing with CI/CD in mind (more on this below)</li>
	<li>Stability and performance improvements</li>
		<ul>
			<li>Implementation of channels and buffering to capacity</li>
			<li>Implementation of caching and auto scaling for clients (In this case React app). Although it's fairly lightcode.</li>
			<li>Implementation of parallelization/concurrency/process waiting lists (workers) for client frontend to increment the capacity to deliver connections to and from backend to more clients at the same time. Also in case huge folders need to be listed, so the request could be implemented in a different manner, paginated, etc. Workers could be goroutines that respond from different channels for every page (etc). For future versions, maybe.</li>
			<li>Implementation of process for workers as a PWA version of the Client to operate and save state even on a connection loss, retrying communication/processes as soon as connection is available</li>
			<li>Implementation of methods and functions that have ability to recover or retry the operation that failed through API and/or CLI: Timers, recovers, etc</li>
			<li>Further testing should be implemented: FastSwitchSli, HandlePath, and other functions have no test implemented yet</li>
			<li>Further Benchmarking tests should be implemented. This should be done to proper evaluate performance. Specially for scaling up and remote using this program.</li>
			<li>API functions have no unit tests. That should be sorted</li>
			<sub>(See <a href="https://golang.org/pkg/testing/">GO Testing</a>)</sub>
			<li>Some more error handling in code could be added, specially with the use of Recover for increased stability</li>
			<sub>It's fairly managed, but some user entering parameters in some cases might develop in error during runtime. I discover some of recovers from ROUTER and implemented solutions, but new errors may appear. Hence the need for more testing</sub>
			<li>Human error workarounds</li>
			<sub>The ability to work even when some minor mistakes are entered when calling the program (via API or CLI). Detection of different quotation symbols, lower/upper casing letters in the middle, missing some symbols etc</sub>
			<li>Slashes before and after paths was one of these. I created a solution for this, but more engineering should be performed.</li>
			<li>Double check and review that code closes every open file or path to improve memory management</li>
		</ul>
	<li>Complexity and Coverage</li>
		<ul>
			<li>Middleware implementation. Specially for cybersecurity purposes, but also for error handling, re-routing, and modularity in different microservices of some parts.</li>
			<li>More argument complexity for CLI and API (Adding flag capacity with flag package)</li>
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
			<li>Filtering capacity (Client program and flags for CLI)</li>
			<sub>Full example could be: blit -path="/usr/local/go" -size="asc" -filter="*.go"</sub>
		</ul>
	<li>Documentation:</li>
		<ul>
			<li>Examples should be added in blit_test.go to be displayed in Blit documentation (godoc)</li>
			<li>Modify testing to fullfill beter the CI/CD tool and maybe OpenAPI aswell, to be integrated with.</li>
		</ul>
	<li>Simplification of functions:</li>
		<ul>
			<li>EncapData() function from blit_cli package is too large. E.g.: Too many parameters returned.</li>
			<li>Code could be fragmented even more to specialize it to make this more modular and reusable. Although the rest of the code it's fairly optimized</li>
			<li>Modularization</li>
			<sub>Subdividing code for different variations of parameters being passed into the program could lead to a faster minor programs that could be run separately, or to fulfill other program's needs. That would lead to a better integration and even completion of a full REST API architectural solution</sub>
		</ul>
	<li>Further Beautification:</li>
		<ul>
			<li>Colouring Folders instead of column specifying whenever a "file" is in reality a folder for CLI</li>
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
		<li>API tokenization and security process to handle new operations depending on tokens and "levels"</li>
		<li>Implementation of Secrets (pub key, priv key, ...) to exchange information between frontend and backend. Encryption, etc</li>
		<sub>Better encapsulation and isolation of data. Migrate from full Capitalized funcs and vars, to a more secure version, maybe a fork, to fullfill some "client" programs needs: Banks, cybersecurity?</sub>
		<li>Database implementation for API conectivity. That could forbid or allow access to data. Private and/or internal API inside a more public API version, for example.</li>
		<ul>
			<li>i.e: User auth, checking credentials, etc</li>
		</ul>
	</ul>
</ul>


