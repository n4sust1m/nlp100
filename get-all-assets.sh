mkdir ./assets

# 01x
curl http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt --output "./assets/hightemp.txt"

# 02x
curl https://www.cl.ecei.tohoku.ac.jp/nlp100/data/jawiki-country.json.gz --output "./assets/jawiki-country.json.gz"

# 04x
curl http://www.cl.ecei.tohoku.ac.jp/nlp100/data/neko.txt --output "./assets/neko.txt"
mecab ./assets/neko.txt -o ./assets/neko.txt.mecab

# unzip files
cd ./assets
gunzip jawiki-country.json.gz