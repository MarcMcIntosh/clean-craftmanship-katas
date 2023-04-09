import Game from "../"


describe('bowling game', () => {
  let game = new Game()

  beforeEach(() => {
    game = new Game()
  })

  function rollMany(times: number, pins: number) {
    for(let i = 0; i < times; i += 1) {
      game.roll(pins)
    }
  }

  function rollSpare() {
    rollMany(2, 5)
  }

  function rollStrike() {
    rollMany(1, 10)
  }


  test('if this does not run then there is an issue with the setup (contact us)', () => {
    expect(true).toBeTruthy()
  })

  test('it can start a new game', () => {
    expect(game).toBeDefined()
  })

  test('can roll', () => {
    game.roll(0)
  })

  test("gutter game", () => {
    rollMany(20, 0)
    expect(game.getScore()).toEqual(0)
  })

  test("all ones", () => {
    rollMany(20, 1)
    expect(game.getScore()).toEqual(20)
  })

  test('one spare', () => {
    rollSpare()
    game.roll(7)
    rollMany(17, 0)
    expect(game.getScore()).toEqual(24)
  })

  test("one strike", () => {
    rollStrike()
    game.roll(2)
    game.roll(3)
    rollMany(16, 0)
    expect(game.getScore()).toEqual(20)
  })

  test('perfect game', () => {
    rollMany(12, 10)
    expect(game.getScore()).toEqual(300)
  })
})