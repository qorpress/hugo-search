# hugo-search [![Build Status](https://travis-ci.org/tischda/hugo-search.svg?branch=master)](https://travis-ci.org/tischda/hugo-search)

A [Bleve](http://www.blevesearch.com) search server for your [Hugo](http://gohugo.io) site.

### Install

~~~
go get github.com/stretchr/testify
go get github.com/kardianos/govendor

git clone https://github.com/tischda/hugo-search

cd hugo-search

govendor sync
make install
~~~


### Dependencies

Note that blevesearch vendors with [gvt](https://github.com/FiloSottile/gvt) and hugo with 
[govendor](https://github.com/kardianos/govendor). I copy-pasted `vendor.json` from hugo and
converted the manifest from bleve manually. Finally, I pinned the versions for bleve and hugo:

~~~
./gvt-manifest-to-govendor.sh > ./vendor-bleve.sh
./vendor-bleve.sh

govendor fetch github.com/blevesearch/bleve/...@v0.5
govendor fetch github.com/spf13/hugo/...@v0.18
govendor fetch github.com/rs/cors
~~~

I agree, this is cumbersome.


### Usage

~~~
Usage of hugo-search:
  -addr string
        http listen address (default ":8080")
  -hugoPath string
        path of the hugo site (default ".")
  -indexPath string
        path of the bleve index (default "indexes/search.bleve")
  -stepAnalysis
        display memory and timing of different steps of the program
  -verbose    verbose output
  -version
        print version and exit
~~~

### Query index

~~~
$ curl http://localhost:8080/api/search.bleve/_search -d '{"query":{"query":"lorem"}}'
{"request":{"query":{"query":"lorem","boost":1},"size":0,"from":0,"highlight":null,"fields":null,"facets":null,"explain":false},"hits":[],"total_hits":2,"max_score":0.15713484143442302,"took":0,"facets":{}}

{"status":{"total":1,"failed":0,"successful":1},"request":{"query":{"query":"lorem","boost":1},"size":0,"from":0,"highlight":null,"fields":null,"facets":null,"explain":false},"hits":[],"total_hits":3,"max_score":0.15713484143442302,"took":0,"facets":{}}
~~~

### Explore index with bleve-explorer

Warning: Cannot use while `hugo-search` is running.

~~~
go get github.com/blevesearch/bleve-explorer

bleve-explorer -dataDir indexes
~~~

check on [http://localhost:8095/](http://localhost:8095/)
