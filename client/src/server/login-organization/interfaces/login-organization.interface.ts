export interface ILoginOrganizationIn {
  document: string | null;
  password: string | null;
}

export interface ILoginOrganizationOut {
  error: boolean;
  accessToken: string;
}
