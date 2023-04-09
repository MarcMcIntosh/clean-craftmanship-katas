class Game {

  private currentRoll: number = 0;
  private rolls: number[] = [];

  private isSpare(frameIndex: number): boolean {

    return this.rolls[frameIndex] + this.rolls[frameIndex + 1] === 10

  }

  private isStrike(frameIndex: number): boolean {

    return this.rolls[frameIndex] === 10

  } 

  roll(pins: number): void {

    this.rolls[this.currentRoll++] = pins;

  }

  getScore(): number {

    let score = 0;

    let frameIndex = 0;

    for(let frame = 0; frame < 10; frame += 1) {

      if(this.isStrike(frameIndex)) {

        score += 10 + this.rolls[frameIndex + 1] + this.rolls[frameIndex + 2]

        frameIndex += 1 

      } else if(this.isSpare(frameIndex)) {

        score += 10 + this.rolls[frameIndex + 2]

        frameIndex += 2

      } else {

        score += this.rolls[frameIndex] + this.rolls[frameIndex + 1]
        
        frameIndex += 2

      }
    }

    return score
  }
}

export default Game