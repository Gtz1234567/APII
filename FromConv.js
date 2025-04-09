const romanNumerals = [
  { roman: 'M', value: 1000 },
  { roman: 'CM', value: 900 },
  { roman: 'D', value: 500 },
  { roman: 'CD', value: 400 },
  { roman: 'C', value: 100 },
  { roman: 'XC', value: 90 },
  { roman: 'L', value: 50 },
  { roman: 'XL', value: 40 },
  { roman: 'X', value: 10 },
  { roman: 'IX', value: 9 },
  { roman: 'V', value: 5 },
  { roman: 'IV', value: 4 },
  { roman: 'I', value: 1 },
];

function fromRoman(roman) {
  let i = 0, value = 0;
  roman = roman.toUpperCase();
  while (i < roman.length) {
    let two = roman.substr(i, 2);
    let one = roman[i];
    let found = romanNumerals.find(r => r.roman === two || r.roman === one);
    if (found) {
      value += found.value;
      i += found.roman.length;
    } else {
      return null;
    }
  }
  return value;
}

export default function handler(req, res) {
  const roman = req.query.roman;
  const result = fromRoman(roman);
  if (!result) {
    return res.status(400).json({ error: 'Número romano inválido' });
  }
  res.status(200).json({ number: result });
}
