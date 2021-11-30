export default function authHeader() {
  let user = JSON.parse(localStorage.getItem('user'));

  if (user.jwt) {
    return { 'jwt-token': user.jwt}; 
  } else {
    return {};
  }
}
