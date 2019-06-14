const str = 'stressed'.split('')

const reversedStr = str.reduce((summary = '', char) => char + summary)

console.log(reversedStr)