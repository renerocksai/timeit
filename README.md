# timeit

A simple little command line utility, to enable command execution time on the Windows command prompt.

Only a few lines of GO, no fancy stuff.

## Download and use it

Best put it somewehere in your PATH after download -- or add the dir you downloaded it to into your PATH.

Then, to use it, just prepend any command with the timeit command; for example:

Execution **without** timeit:

```bash
λ dir
 Datenträger in Laufwerk C: ist Windows
 Volumeseriennummer: 4C97-E486

 Verzeichnis von C:\Users\rene.schallner\go\src\github.com\renerocksai\timeit

17.04.2019  18:31    <DIR>          .
17.04.2019  18:31    <DIR>          ..
17.04.2019  18:16                22 .gitignore
17.04.2019  18:31               587 main.go
17.04.2019  18:24               362 README.md
17.04.2019  18:31         2.453.504 timeit.exe
               4 Datei(en),      2.454.475 Bytes
               2 Verzeichnis(se), 160.585.662.464 Bytes frei

```

Execution **with** timeit:

```bash
λ timeit dir
Elapsed: 33.1695ms
 Datenträger in Laufwerk C: ist Windows
 Volumeseriennummer: 4C97-E486

 Verzeichnis von C:\Users\rene.schallner\go\src\github.com\renerocksai\timeit

17.04.2019  18:31    <DIR>          .
17.04.2019  18:31    <DIR>          ..
17.04.2019  18:16                22 .gitignore
17.04.2019  18:31               587 main.go
17.04.2019  18:24               362 README.md
17.04.2019  18:31         2.453.504 timeit.exe
               4 Datei(en),      2.454.475 Bytes
               2 Verzeichnis(se), 160.585.322.496 Bytes frei
```


## Build it yourself

Provided you have go:

```bash
go get github.com/renerocksai/timeit
go install github.com/renerocksai/timeit   # if you're not in the project directory
# go install    # if you are
```

