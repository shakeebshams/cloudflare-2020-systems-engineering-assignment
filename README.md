# Systems Engineering Requests Tool

## What is it?

I built a CLI tool for Cloudflare's 2021 Engineering Assesment which makes HTTP requests using sockets directly rather than using a library. The program can also return request statistics such as fastest/slowest times, median/mean times, response sizes etc.


### Installation
Install go from here https://golang.org/doc/install.
Once you have go installed, you should be ready to go!

### Usage

* Run ```go run . --help ``` to open up a help prompt with all usages and different flags allowed

* Run ```go run . --url <url>``` to get a full response from a website including status code and html

* Run ```go run . --profile <url> <number of requests>``` to profile a specific number of requests to a url

## Results

### General Results

Cloudflare Worker /links site
![workerlinkssite](screenshots/linksGen.png)

### Profile Stats Results
Cloudflare Worker Site
![workersite](screenshots/workerProfile.png)

Cloudflare Worker /links site
![workersite](screenshots/linksProfile.png)

Cloudflare.com
![workersite](screenshots/cloudflareProfile.png)

Cycling Stats site
![cycling ranking profile](screenshots/cyclingProfile.png)

Google.com Profile
![Google Profile](screenshots/googleProfile.png)

### Analysis of Profile Stats Results


