jest.mock('axios');
const axios = require('axios');
const {getCurrentQuestion, postQuestion} = require("./serverUtils");



describe('getCurrentQuestion', () => {
  it('should send an axios GET request to the correct endpoint', () => {
    let actualUrl;
    axios.get.mockImplementation((urlParam) => {
      return new Promise((resolve) => {
        actualUrl = urlParam;
        resolve({data: null});
      });
    });
    let expectedUrl = `/api/question/current`;
    getCurrentQuestion().then(()=>{
      expect(actualUrl).toBe(expectedUrl);
    });
  });
});

describe('postQuestion', () => {
  it('should send an axios POST request to the correct endpoint', () => {
    let actualUrl, actualData;
    axios.post.mockImplementation((url, data) => {
      return new Promise((resolve) => {
        actualUrl = url;
        actualData = data;
        resolve({data: null});
      });
    });
    const inputQuestion = 'Input Question';
    const expectedUrl = '/api/question'
    postQuestion(inputQuestion).then(() => {
      expect(actualUrl).toBe(expectedUrl)
    })
  })
})