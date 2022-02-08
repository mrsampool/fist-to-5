import axios from "axios";
import {flushPromises, mount} from '@vue/test-utils';
import StudentView from "../StudentView.vue";

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

describe('Student View', () => {

  it('should GET current question & render it', async() => {
    const wrapper = mount(StudentView);
    await flushPromises()
    expect(axios.get).toHaveBeenCalledTimes(1);
    expect(axios.get).toHaveBeenCalledWith('/api/question/current');
    expect(wrapper.html()).toContain(mockedCurrentQuestion.data.text)
  });

  it('should POST user\'s response', async() => {
    for (const value of [0, 1, 2, 3, 4, 5]){
      const wrapper = mount(StudentView);
      await flushPromises();
      await wrapper.get(`#radio-${value}`).setValue(true);
      await wrapper.get('form').trigger('submit');
      expect(mockedPost.value).toBe(value);
    };
    expect(axios.post).toHaveBeenCalledTimes(6);
  });

});