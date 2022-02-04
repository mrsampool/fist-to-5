jest.mock('axios');
const axios = require('axios');
const {getCurrentQuestion} = require("./serverUtils");

axios.post
  .mockImplementation((url, data) => new Promise((resolve) => resolve({ url, data })));

describe('getCurrentQuestion', () => {
  it('should send an axios GET request to the correct endpoint', () => {
    let actualUrl;
    let expectedUrl = `/api/question/current`;
    axios.get.mockImplementation((urlParam) => {
      return new Promise((resolve) => {
        actualUrl = urlParam;
        resolve();
      });
    });
    getCurrentQuestion().then(()=>{
      expect(actualUrl).toBe(expectedUrl);
    });
  });
});

describe('postQuestion', () => {
  it('should send a POST request to the current endpoint', () => {
    let actualUrl;
    let expectedUrl = `/api/question/current`;
    axios.get.mockImplementation((urlParam) => {
      return new Promise((resolve) => {
        expectedUrl = urlParam;
        resolve();
      });
    });
    getCurrentQuestion().then(()=>{
      expect(actualUrl).toBe(expectedUrl);
    });
  })
})