type Inputs = string | number

const generatePhrase = (x: Inputs, y: Inputs, z: Inputs): string =>
  `${x}時の${y}は${z}`

console.log(generatePhrase(12, '気温', 22.4))