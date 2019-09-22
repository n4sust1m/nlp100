const str: string = 'I couldn\'t believe that I could actually understand what I was reading : the phenomenal power of the human mind .'

const dividedStrs: Array<string> = str.split(' ')

const typoglycemia: Array<string> = dividedStrs.map((word) => {
  if(word.length <= 4)
    return word

  let charactors: Array<string> = word.split('')

  // fisher yates algorhythm ; O(n) 
  for(let i = 1; i < charactors.length - 1; i++){
    let r = 1 + Math.floor(Math.random() * i)
    let tmp = charactors[i]
    charactors[i] = charactors[r]
    charactors[r] = tmp
  }
  return charactors.reduce((accum, currentValue) => accum + currentValue)
})

console.log(typoglycemia.reduce((accum, currentValue) => accum + ' ' + currentValue))