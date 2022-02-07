import axios from "axios";

export function getCurrentQuestion(){
  return new Promise((resolve, reject) => {
    axios.get(`/api/question/current`)
      .then(({data}) => resolve(data))
      .catch(err => reject(err));
  });
}

export function postQuestion(text){
  return new Promise((resolve, reject) => {
    if (text) {
      axios.post(`/api/question`, { text })
        .then(() => resolve())
        .catch(err => reject(err))
    }
  });
}

export function getCurrentResponses(){
  return new Promise((resolve, reject) => {
    axios.get(`/api/responses/current`)
      .then(({data}) => resolve(data))
      .catch(err => reject(err));
  });
}

export function postResponse(currQuestion, currStudent, value) {
  return new Promise((resolve, reject) => {
    if (currQuestion && currQuestion.id && value !== null){
      const resData = {
        studentId: currStudent.id,
        value: value,
      }
      axios.post(`/api/responses/${currQuestion.id}`, resData)
        .then(() => resolve())
        .catch(err => reject(err))
    }
  });
}