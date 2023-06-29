export interface IRegisterOrganizationIn {
  name: string | null;
  password: string | null;
  logo?: string | null;
  document: string | null;
}

export interface IRegisterOrganizationOut {
  error: boolean;
  message: string;
  infos?: object;
}
