#! /bin/bash

wget -O ./res/jawiki-country.json.gz http://www.cl.ecei.tohoku.ac.jp/nlp100/data/jawiki-country.json.gz &&
  gzip -d ./res/jawiki-country.json.gz