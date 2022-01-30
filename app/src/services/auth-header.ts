import { UserAuthHeader, UserStore } from "@/types/user";

export default function authHeader():UserAuthHeader {
  const authHeader: UserAuthHeader= {} as UserAuthHeader;
  const rawUser = localStorage.getItem('user') ? localStorage.getItem('user') : '{}';

  if (typeof rawUser !== 'string') return authHeader;

  const user: UserStore = JSON.parse(rawUser);
  authHeader["x-access-token"]= user.jwt;
  return authHeader;
}
