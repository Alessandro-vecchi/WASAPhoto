<a name="readme-top"></a>
<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h2>WASAPhoto</h2>

  
</div>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#demo">Demo</a></li>
    <li>
      <a href="#project-structure">Project Structure</a>
      <ul>
        <li><a href="#Go-vendoring">Go vendoring</a></li>
        <li><a href="#Node/NPM-vendoring">Node/NPM vendoring</a></li>
        <li><a href="#How-to-set-up-a-new-project-from-this-template">How to set up a new project from this template</a></li>
        <li><a href="#How-to-build">How to build</a></li>
        <li><a href="#How-to-run">How to run (in development mode)</a></li>
      </ul>
    </li>
    <li><a href="#known-issues">Known Issues</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! 
Directly from your PC, you can upload your photos, and they will be visible to everyone who is following you.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [![Swagger][Swagger-badge]][Swagger-url]
* [![Go][Go-badge]][Go-url]
* [![SQLite][SQLite-badge]][SQLite-url]
* [![HTML5][HTML-badge]][HTML-url]
* [![CSS3][CSS-badge]][CSS-url]
* [![JavaScript][JavaScript-badge]][JavaScript-url]
* [![Vue.js][Vue-badge]][Vue-url]
* [![Docker][Docker-badge]][Docker-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Demo -->
### Demo

Here you can see a demo of the app:

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Project structure

* `cmd/` contains all executables; Go programs here should only do "executable-stuff", like reading options from the CLI/env, etc.
	* `cmd/healthcheck` is an example of a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains an example of a web API server daemon
* `demo/` contains a demo config file
* `doc/` contains the documentation (an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains an example of an API server
	* `service/globaltime` contains a wrapper package for `time.Time` (useful in unit testing)
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` is an example of a web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding

Other project files include:
* `open-npm.sh` starts a new (temporary) container using `node:lts` image for safe web frontend development

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Go vendoring

This project uses [Go Vendoring](https://go.dev/ref/mod#vendoring). You must use `go mod vendor` after changing some dependency (`go get` or `go mod tidy`) and add all files under `vendor/` directory in your commit.

For more information about vendoring:

* https://go.dev/ref/mod#vendoring
* https://www.ardanlabs.com/blog/2020/04/modules-06-vendoring.html

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Node/NPM vendoring

This repository contains the `webui/node_modules` directory with all dependencies for Vue.JS. You should commit the content of that directory and both `package.json` and `package-lock.json`.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## How to set up a new project from this template

You need to:

* Change the Go module path to your module path in `go.mod`, `go.sum`, and in `*.go` files around the project
* Rewrite the API documentation `doc/api.yaml`
* If no web frontend is expected, remove `webui` and `cmd/webapi/register-webui.go`
* If no cronjobs or health checks are needed, remove them from `cmd/`
* Update top/package comment inside `cmd/webapi/main.go` to reflect the actual project usage, goal, and general info
* Update the code in `run()` function (`cmd/webapi/main.go`) to connect to databases or external resources
* Write API code inside `service/api`, and create any further package inside `service/` (or subdirectories)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## How to run (in development mode)

You can launch the backend only using:

```shell
go run ./cmd/webapi/
```

If you want to launch the WebUI, open a new tab and launch:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Known issues

### Apple M1 / ARM: `failed to load config from`...

If you use Apple M1/M2 hardware, or other ARM CPUs, you may encounter an error message saying that `esbuild` (or some other tool) has been built for another platform.

If so, you can fix issuing these commands **only the first time**:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm install
exit
# Now you can continue as indicated in "How to build/run"
```

**Use these instructions only if you get an error. Do not use it if your build is OK**.

### My build works when I use `npm run dev`, however there is a Javascript crash in production/grading

Some errors in the code are somehow not shown in `vite` development mode. To preview the code that will be used in production/grading settings, use the following commands:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-prod
npm run preview
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [x] Add Badges
- [x] Add back to top links
- [x] Add Video
- [x] Add License 
- [x] Add Contact
- [ ] Add Usage
- [ ] Add Getting Started
- [ ] Add Changelog
      
<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Alessandro Vecchi - alessandro.vecchi66@gmail.com

Project Link: [https://github.com/Alessandro-vecchi/WASAPhoto](https://github.com/Alessandro-vecchi/WASAPhoto)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/Alessandro-vecchi/WASAPhoto.svg?style=for-the-badge
[contributors-url]: https://github.com/Alessandro-vecchi/WASAPhoto/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/Alessandro-vecchi/WASAPhoto.svg?style=for-the-badge
[forks-url]: https://github.com/Alessandro-vecchi/WASAPhoto/network/members

[stars-shield]: https://img.shields.io/github/stars/Alessandro-vecchi/WASAPhoto.svg?style=for-the-badge
[stars-url]: https://github.com/Alessandro-vecchi/WASAPhoto/stargazers

[issues-shield]: https://img.shields.io/github/issues/Alessandro-vecchi/WASAPhoto.svg?style=for-the-badge
[issues-url]: https://github.com/Alessandro-vecchi/WASAPhoto/issues

[license-shield]: https://img.shields.io/github/license/Alessandro-vecchi/WASAPhoto.svg?style=for-the-badge
[license-url]: https://github.com/Alessandro-vecchi/WASAPhoto/blob/master/LICENSE.txt

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/alessandro-v-6711

[product-screenshot]: images/screenshot.png

[Swagger-badge]: https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white
[Swagger-url]: https://swagger.io/

[Go-badge]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/

[SQLite-badge]: https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white
[SQLite-url]: https://sqlite.org/

[HTML-badge]: https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white
[HTML-url]: https://html.com/

[CSS-badge]: https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white
[CSS-url]: https://CSS.com/

[JavaScript-badge]: https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E
[JavaScript-url]: https://JavaScript.com/


[Vue-badge]: https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D
[Vue-url]: https://JavaScript.com/

[Docker-badge]: https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white
[Docker-url]: https://Docker.com/

