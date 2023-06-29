import { ServerAdapter } from "../adapter/server-adapter";
import {
  IRegisterOrganizationIn,
  IRegisterOrganizationOut,
} from "./interfaces/register-organization.interface";
import Api from "../api-instance";
class RegisterOrganizationServer extends ServerAdapter<
  IRegisterOrganizationIn,
  Promise<IRegisterOrganizationOut>
> {
  async Run(data: IRegisterOrganizationIn): Promise<IRegisterOrganizationOut> {
    const response = await Api.post("/organization", {
      Name: data.name,
      Document: data.document,
      Password: data.password,
      Logo: data.logo,
    });
    console.log(response);
    return this.parserResponse(response);
  }
}

export default new RegisterOrganizationServer();
