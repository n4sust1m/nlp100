const str = 'Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can'.split(' ')

const oneCharIndex = [1, 5, 6, 7, 8, 9, 15, 19]
const elements = str.map((v, i) => oneCharIndex.find(v => v - 1 === i) ? v.slice(0, 1) : v.slice(0, 2))

console.log(elements)