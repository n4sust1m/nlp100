type BiGramGen = ((string) => Array<string>)

const getWordBiGram: BiGramGen = (str) => str.split(' ').reduce((p = new Array(), c, i) => i % 2 == 1 ? ((p, c) => {p.append(c); return p})(p, c) : ((p, c) => {p[p.length - 1] += ' ' + c; return p})(p, c))
const getCharBiGram: BiGramGen = (str) => str.reduce((p = new Array(),c,i) => i % 2 === 1 ? ((p, c) => {p.append(c); return p})(p, c) : ((p, c) => {p[p.length - 1] += ' ' + c; return p})(p, c))

const getNGram: ((string) => any) = (str) => {
    console.log(getWordBiGram(str))
    console.log(getCharBiGram(str))
}

getNGram('I am an NLPer')