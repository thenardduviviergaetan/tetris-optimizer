<div align="center" id="top"> 
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/50/All_5_free_tetrominoes.svg/800px-All_5_free_tetrominoes.svg.png" width="250" height="150" alt="tetris" />

  &#xa0;

  <!-- <a href="https://versionfinal.netlify.app">Demo</a> -->
</div>

<h1 align="center">tetris-optimiser</h1>

<p align="center">
   <img src="https://img.shields.io/badge/language-GO-green?style=flat" alt="language" />
</p>

<!-- Status -->

<!-- <h4 align="center"> 
	🚧  Versionfinal 🚀 Under construction...  🚧
</h4> 

<hr> -->

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-author">Author</a> &#xa0;
</p>

<br>

## :dart: About ##

It's an optimized Tetris solver that takes a txt file as input to retrieve the Tetrominoes. 

## :rocket: Technologies ##

The following tools were used in this project:

- [Git](https://git-scm.com)
- [GO](https://go.dev/doc/install) 
- [VScode](https://code.visualstudio.com/)

## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com) , [GO](https://go.dev/doc/install) and [VScode](https://code.visualstudio.com/) installed.

## :checkered_flag: Starting ##

```bash
# Clone this project
$ git clone https://zone01normandie.org/git/gthenard/tetris-optimizer.git

# Access
$ cd tetris-optimizer

# Run the project
$ go run . <datadeteste>

```
### Exemple ###
``` 
## Usage

```console
$ cat -e sample.txt
...#$
...#$
...#$
...#$
$
....$
....$
....$
####$
$
.###$
...#$
....$
....$
$
....$
..##$
.##.$
....$
$
....$
.##.$
.##.$
....$
$
....$
....$
##..$
.##.$
$
##..$
.#..$
.#..$
....$
$
....$
###.$
.#..$
....$
$ go run . sample.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
$
```

## :memo: Author ##

Made with by <a href="https://zone01normandie.org/git/gthenard" target="_blank">gthenard</a>

&#xa0;

<a href="#top">Back to top</a>
