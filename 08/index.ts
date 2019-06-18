const cipher: (string) => string = (planeText) =>
  planeText.split('')
    .map((v: string) => 0x61 <= v.charCodeAt(0) && v.charCodeAt(0) <= 0x7a ? String.fromCharCode(219 - v.charCodeAt(0)) : v)
    .join('')

console.log(cipher('hello'))
console.log(cipher('konnadsfa6#$%^sdad'))
console.log(cipher('こんだだ１'))
console.log(cipher('hel4'))