const localstorage = require('./locastorage')
const url = 'http://127.0.0.1:3000'

export function login (email, password) {
  return new Promise((resolve, reject) => {
    fetch(url + '/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: email,
        password: password
      })
    })
      .catch(err => {
        reject(err)
      })
      .then(response => response.json())
      .then(data => {
        localstorage.setJwt(data.jwt)
        resolve()
      })
  })
}
