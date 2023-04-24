# SDE Code Exercise

## Initial Setup

<details>
  <summary>Click to expand!</summary>
  
  1. You need to install `Go` on your computer in order to run the code and the tests. Follow the [instructions](https://go.dev/dl/) for your current platform

  2. Add `GOPATH` and `GOBIN` to your environment variables

```bash
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/bin"
```

</details>


## Run the code

  1. In order to the code, you need to provide two flags to tell the program where to find the `destinations` file and the `drivers` file.

    * `destinations` file have a list of street names, separated by carriage return, i.e.

```
street name #1
street name #2
street name #3
```

    * `drivers` file have a list of drivers names, separated by carriage return, i.e.
```
john doe
bob martins
```

  2. To run the code with the current files on the repo, execute: 

```bash
go run . \
  -destinations destinations.txt
  -drivers drivers.txt
```


  3. The output will be ordered by Suitability Score (SS), 

```
{SS:13.5 Driver:Daniel Martinez Dest:Cerrada Otono #15, 04102}
{SS:9 Driver:Jason Statham Dest:First Avenue Street #459, Nashville, Tennessee}
{SS:7.5 Driver:Alan Moore Dest:Nueva Santa Maria, Azcapotzalco, 02800 Ciudad de Mexico}
```
