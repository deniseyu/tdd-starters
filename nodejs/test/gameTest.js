const expect = require('chai').expect
const game = require('../game.js')

describe('Game of life', () => {
  it('should say yo', () => {
    expect(game.doSomething()).to.equal('yo')
  })
})
