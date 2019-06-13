const str = 'パタトクカシーー'.split('')

const processedStr = str.reduce((p = '', c, i) => i % 2 == 0 ? p + c : p)

console.log(processedStr)