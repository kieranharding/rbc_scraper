`rbc_scraper`
=============

A command line tool which walks all the RBC Statement PDF files in the current directory and parses all the transactions into a `transactions.csv` file.

The tool has been tested with *Checking* and *Visa* statements so far.  The tool has been tested on Mac OSX, but it should work on Linux as well without modification.  If you would like to make this work with Windows, it would require slight modification.


Dependencies
------------

`rbc_scraper` depends on the [`pdf2htmlEX`](https://github.com/coolwanglu/pdf2htmlEX/wiki/Download) command line tool for converting the PDF to an HTML in order for the data to be parsed out.


Usage
-----

**Downloading Binary**
```bash
$ /path/to/dir/with/rbc/statements
$ wget -O rbc_scraper https://github.com/swill/rbc_scraper/raw/master/bin/rbc_scraper_darwin_amd64
$ chmod +x rbc_scraper
$ ./rbc_scraper
```

**Building Locally**
```bash
$ git clone https://github.com/swill/rbc_scraper.git
$ cd rbc_scraper
$ go install
$ cd /path/to/dir/with/rbc/statements
$ rbc_scraper
$ cat transactions.csv
```