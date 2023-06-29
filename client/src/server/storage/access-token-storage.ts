const NAME = "@access_toke";

export const AccessTokenStorage = {
  get: () => {
    return localStorage.getItem(NAME);
  },
  post: (token: string) => {
    localStorage.setItem(NAME, token);
  },
};
