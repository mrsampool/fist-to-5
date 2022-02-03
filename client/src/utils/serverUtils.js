import axios from "axios";

export function getCurrentQuestion(component){
  return new Promise((resolve, reject) => {
    axios.get(`/api/question/current`)
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
  })
}