const str1 = 'paraparaparadise'
const str2 = 'paragraph'

function getBiGram (str: string): Array<string> {
  const separatedStr: string[] = str.split('')
  let biGramedStr: string[] = []
  for(let i: number = 0; separatedStr.length > 0; i++){
    if(i % 2 === 0)
      biGramedStr.push(separatedStr.shift())
    else
      biGramedStr[biGramedStr.length - 1] += separatedStr.shift()
  }
  return biGramedStr
}

const x: Array<string> = getBiGram(str1)
const y: Array<string> = getBiGram(str2)

const getAnd = (str1: Array<string>, str2: Array<string>): Array<String> => 
  Array.from(new Set(str1))
    .filter((v1) => str2.find((v2) => v1 === v2))

const getOr = (str1: Array<string>, str2: Array<string>): Array<String> => 
  Array.from(new Set(str1.concat(str2)))

const getSa = (str1: Array<string>, str2: Array<string>): Array<String> => 
  getOr(str1, str2)
    .filter((v) => getAnd(str1, str2).indexOf(v) === -1)

const search_se = (str: Array<string>): boolean =>
  str.indexOf('se') !== -1

console.log(x)
console.log(y)
console.log(getAnd(x, y))
console.log(getOr(x, y))
console.log(getSa(x, y))

console.log(search_se(x))
console.log(search_se(y))