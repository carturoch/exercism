const ALPHA = 'abcdefgihjklmnopqrstuvwxyz';

class Pangram {
  constructor(phrase) {
    this.phrase = phrase;
  }

  isPangram() {
    if (this.phrase.length < ALPHA.length) {
      return false;
    }
    
    return ALPHA
      .split('')
      .every((c) => this.phrase.toLowerCase().includes(c));
  }
}

module.exports = Pangram;