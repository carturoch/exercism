const ALPHA = 'abcdefgihjklmnopqrstuvwxyz'; 

class Pangram {
  constructor(phrase) {
    this.phrase = phrase;
  }

  isPangram() {
    if (this.phrase.length < ALPHA.length) {
      return false;
    }
    let map = {};
    ALPHA
      .split('')
      .forEach((letter) => map[letter] = 0);
    
    this
      .phrase
      .toLowerCase()
      .split('')
      .filter((c) => ALPHA.includes(c))
      .forEach((c) => map[c]++)

    return Object.values(map).every((c) => c > 0);
  }
}

module.exports = Pangram;