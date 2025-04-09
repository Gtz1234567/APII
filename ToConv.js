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

function toRoman(num) {
  let result = '';
  for (let { roman, value } of romanNumerals) {
    while (num >= value) {
      result += roman;
      num -= value;
    }
  }
  return result;
}

export default function handler(req, res) {
  const number = parseInt(req.query.number);
  if (isNaN(number) || number < 1 || number > 3999) {
    return res.status(400).json({ error: 'Número inválido (1-3999)' });
  }
  res.status(200).json({ roman: toRoman(number) });
}
