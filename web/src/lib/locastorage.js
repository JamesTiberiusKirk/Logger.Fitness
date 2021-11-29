export function setJwt (jwt) {
  localStorage.setItem('jwt-token', jwt)
}

export function getJwt () {
  return localStorage.getItem('jwt-token')
}
