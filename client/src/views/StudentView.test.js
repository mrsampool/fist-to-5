const axios = require("axios");
const {getCurrentQuestion} = require("./serverUtils");

describe('Student View', () => {
  it('should request the current question and render it to the screen', () => {
    let actualUrl;
    axios.get.mockImplementation(() => {
      return new Promise((resolve) => {
        resolve({data: null});
      });
    });
    let expectedUrl = `/api/question/current`;
    getCurrentQuestion().then(()=>{
      expect(actualUrl).toBe(expectedUrl);
    });
  });
});