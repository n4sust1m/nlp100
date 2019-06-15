type BiGramGen = ((string) => Array<string>)

const getWordBiGram: BiGramGen = (str) => str.split(' ').reduce((p: Array<string> = new Array(), c, i) => i % 2 == 0 ? ((p, c) => {p.push(c); return p})(p, c) : ((p, c) => [...p, p[p.length - 1] + ' ' + c])(p, c))
const getCharBiGram: BiGramGen = (str) => str.reduce((p: Array<string> = new Array(),c,i) => i % 2 === 0 ? ((p, c) => {p.push(c); return p})(p, c) : ((p, c) => [...p, p[p.length - 1] + c])(p, c))

const getNGram: ((string) => any) = (str) => {
    console.log(getWordBiGram(str).filter((v, i) => i % 2 == 1))
    console.log(getCharBiGram(str.split('')).filter((v, i) => i % 2 == 1))
}

getNGram('I am an NLPer')