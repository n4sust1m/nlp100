const articles = JSON.parse(require('fs').readFileSync("./res/jawiki-country.json"))

const articleAboutBriten = articles
  .filter((v) => !!v.text.match(/イギリス/))

console.log(articleAboutBriten)