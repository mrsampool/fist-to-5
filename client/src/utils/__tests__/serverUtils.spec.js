jest.mock('axios');
const axios = require('axios');
const {getCurrentQuestion, postQuestion} = require("../serverUtils");

const mockedCurrentQuestion = { data: { id: 1, text: 'Test Question' } };
let mockedPost;

jest.mock('axios', () => ({
  get: jest.fn(() => {
    return new Promise((res) => res(mockedCurrentQuestion))
  }),
  post: jest.fn((url, data) => {
    return new Promise((res) => {
      mockedPost = data;
      res();
    })
  }),
}));

describe('getCurrentQuestion', () => {
  it('should send an axios GET request to the correct endpoint', async () => {
    let expectedUrl = `/api/question/current`;
    await getCurrentQuestion();
    expect(axios.get).toHaveBeenCalledTimes(1);
    expect(axios.get).toHaveBeenCalledWith('/api/question/current');
  });
});

describe('postQuestion', () => {
  it('should send a POST request to the current endpoint', async () => {
    const testNewQuestion = { text: 'Test New Question'}
    await postQuestion(testNewQuestion.text);
    expect(axios.post).toHaveBeenCalledTimes(1);
    expect(axios.post).toHaveBeenCalledWith(`/api/question`, testNewQuestion);
  })
})