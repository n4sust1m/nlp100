const str1 = 'パトカー'.split('')
const str2 = 'タクシー'.split('')

const mergedStr = [...(new Array(str1.length + 1))].reduce((p = '', c, i) => p + str1[i-1] + str2[i-1])

console.log(mergedStr)