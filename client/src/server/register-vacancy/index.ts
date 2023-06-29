import { ServerAdapter } from "../adapter/server-adapter";
import {
  IRegisterVacancyIn,
  IRegisterVacancyOUt,
} from "./interfaces/register-vacancy.interfaces";
import Api from "../api-instance";

class RegisterVacancyServer extends ServerAdapter<
  IRegisterVacancyIn,
  Promise<IRegisterVacancyOUt>
> {
  async Run(data: IRegisterVacancyIn): Promise<IRegisterVacancyOUt> {
    console.log(data);
    const response = await Api.post("/vacancy", {
      tilte: data.tilte,
      body: data.body,
      location: Number(data.location),
      level: Number(data.level),
      salary_range: Number(data.salary_range),
      organization_id: data.organization_id,
    });

    return this.parserResponse(response);
  }
}

export default new RegisterVacancyServer();
