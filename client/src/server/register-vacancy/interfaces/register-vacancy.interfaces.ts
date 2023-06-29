export interface IRegisterVacancyIn {
  tilte: string | null;
  body: string | null;
  location: number | null;
  level: number | null;
  salary_range: number | null;
  organization_id: string | null;
}

export interface IRegisterVacancyOUt {
  error: boolean;
  message: string;
  infos?: object;
}
